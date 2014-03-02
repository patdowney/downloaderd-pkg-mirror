package http

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/url"

	"github.com/gorilla/mux"
	"github.com/patdowney/downloaderd-pkg/common"
	"github.com/patdowney/downloaderd-pkg/deb"
	api "github.com/patdowney/downloaderd-pkg/downloaderdapi"
)

type DebianResource struct {
	Clock         common.Clock
	DebianService *deb.DebianService
	router        *mux.Router
}

func NewDebianResource(s *deb.DebianService) *DebianResource {
	return &DebianResource{
		Clock:         &common.RealClock{},
		DebianService: s}
}

func (r *DebianResource) RegisterRoutes(parentRouter *mux.Router) {

	parentRouter.HandleFunc("/packages-handler", r.PackagesHandler()).Methods("POST").Name("packages-handler")
	parentRouter.HandleFunc("/release-handler", r.ReleaseHandler()).Methods("POST").Name("release-handler")
	parentRouter.HandleFunc("/deb-handler", r.DebHandler()).Methods("POST").Name("deb-handler")

	parentRouter.HandleFunc("/default-handler", r.DebHandler()).Methods("POST").Name("default-handler")

	r.router = parentRouter
}

func (r *DebianResource) decodeDownload(body io.Reader) (*api.Download, error) {
	decoder := json.NewDecoder(body)
	var download api.Download
	err := decoder.Decode(&download)

	return &download, err
}

func (r *DebianResource) ReleaseHandler() http.HandlerFunc {
	return func(rw http.ResponseWriter, req *http.Request) {
		download, err := r.decodeDownload(req.Body)
		if err != nil {
			status := http.StatusBadRequest
			http.Error(rw, http.StatusText(status), status)
			log.Print(err)
			return
		}

		log.Printf("release-callback received for: %s", download.URL)
		dataLocation := download.GetLink("data")
		res, err := http.Get(dataLocation)
		if err != nil {
			log.Print(err)
			return
		}
		defer res.Body.Close()

		if res.StatusCode == http.StatusOK {
			releaseData := deb.ParseRelease(res.Body)
			checksumType := "SHA256"
			fileRefs := releaseData.FileReferences(checksumType)

			releaseUrl, err := url.Parse(download.URL)
			if err != nil {
				log.Print(err)
				panic(err)
			}

			for _, fileRef := range fileRefs {
				if fileRef != nil {
					fileRefUrl, _ := releaseUrl.Parse(fileRef.FileName)
					callback := r.DebianService.CallbackForFile(fileRef.FileName)

					fileDownloadRequest := api.Request{
						URL:          fileRefUrl.String(),
						Checksum:     fileRef.Checksum,
						ChecksumType: checksumType,
						Callback:     callback}

					_, err := r.DebianService.Client.RequestDownload(&fileDownloadRequest)
					if err != nil {
						log.Print(err)
					}
				}
			}
		}

		rw.WriteHeader(http.StatusOK)
	}
}

func (r *DebianResource) DefaultHandler() http.HandlerFunc {
	return func(rw http.ResponseWriter, req *http.Request) {
		download, err := r.decodeDownload(req.Body)
		if err != nil {
			status := http.StatusBadRequest
			http.Error(rw, http.StatusText(status), status)
			log.Print(err)
			return
		}
		log.Printf("default-handler processed for: %s", download.URL)
	}
}

func (r *DebianResource) DebHandler() http.HandlerFunc {
	return func(rw http.ResponseWriter, req *http.Request) {
		download, err := r.decodeDownload(req.Body)
		if err != nil {
			status := http.StatusBadRequest
			http.Error(rw, http.StatusText(status), status)
			log.Print(err)
			return
		}
		log.Printf("deb-handler processed for: %s", download.URL)
	}
}

func (r *DebianResource) PackagesHandler() http.HandlerFunc {
	return func(rw http.ResponseWriter, req *http.Request) {
		download, err := r.decodeDownload(req.Body)
		if err != nil {
			status := http.StatusBadRequest
			http.Error(rw, http.StatusText(status), status)
			log.Print(err)
			return
		}

		dataLocation := download.GetLink("data")

		r.DebianService.ProcessPackages(
			"http://archive.ubuntu.com/ubuntu",
			download.URL,
			dataLocation)

		log.Printf("packages-handler processed for: %s", download.URL)

		rw.WriteHeader(http.StatusOK)
	}
}
