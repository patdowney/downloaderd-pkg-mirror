package http

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"

	"github.com/gorilla/mux"
	"github.com/patdowney/downloaderd-deb/common"
	api "github.com/patdowney/downloaderd-deb/downloaderdapi"
	"github.com/patdowney/downloaderd-deb/rpm"
)

type RepomdResource struct {
	Clock         common.Clock
	RepomdService *rpm.RepomdService
	router        *mux.Router
}

func NewRepomdResource(s *rpm.RepomdService) *RepomdResource {
	return &RepomdResource{
		Clock:         &common.RealClock{},
		RepomdService: s}
}

func (r *RepomdResource) RegisterRoutes(parentRouter *mux.Router) {

	parentRouter.HandleFunc("/metadata-handler", r.MetadataHandler()).Methods("POST").Name("metadata-handler")
	parentRouter.HandleFunc("/repomd-handler", r.RepomdHandler()).Methods("POST").Name("repomd-handler")
	parentRouter.HandleFunc("/rpm-handler", r.RPMHandler()).Methods("POST").Name("rpm-handler")

	parentRouter.HandleFunc("/default-handler", r.DefaultHandler()).Methods("POST").Name("default-handler")

	r.router = parentRouter
}

func (r *RepomdResource) decodeDownload(body io.Reader) (*api.Download, error) {
	decoder := json.NewDecoder(body)
	var download api.Download
	err := decoder.Decode(&download)

	return &download, err
}

func (r *RepomdResource) RepomdHandler() http.HandlerFunc {
	return func(rw http.ResponseWriter, req *http.Request) {
		download, err := r.decodeDownload(req.Body)
		if err != nil {
			status := http.StatusBadRequest
			http.Error(rw, http.StatusText(status), status)
			log.Print(err)
			return
		}

		log.Printf("repomd-callback received for: %s", download.URL)
		dataLocation := download.GetLink("data")
		res, err := http.Get(dataLocation)
		if err != nil {
			status := http.StatusInternalServerError
			http.Error(rw, http.StatusText(status), status)
			log.Print(err)
			return
		}
		defer res.Body.Close()

		if res.StatusCode == http.StatusOK {
			log.Print("statusok")
			repomd, err := rpm.ParseRepomd(res.Body)
			if err != nil {
				status := http.StatusInternalServerError
				http.Error(rw, http.StatusText(status), status)
				log.Print(err)
				return
			}

			repomdURL, err := url.Parse(download.URL)
			if err != nil {
				status := http.StatusInternalServerError
				http.Error(rw, http.StatusText(status), status)
				log.Print(err)
				return
			}

			//checksumType := "sha256"
			//fileRefs := repomd.FileReferences(checksumType)

			for _, dataSection := range repomd.Data {
				dataHref := dataSection.Location.Href
				dataURL, _ := repomdURL.Parse(fmt.Sprintf("../%s", dataHref))
				log.Printf("dataURL: %s", dataURL.String())
				callback := r.RepomdService.CallbackForFile(dataHref)

				fileDownloadRequest := api.Request{
					URL:          dataURL.String(),
					Checksum:     dataSection.Checksum.Value,
					ChecksumType: dataSection.Checksum.Type,
					Callback:     callback}

				_, err := r.RepomdService.Client.RequestDownload(&fileDownloadRequest)
				if err != nil {
					log.Print(err)
				}

			}
			rw.WriteHeader(http.StatusOK)
		} else {
			log.Print("statusnotok")
			rw.WriteHeader(http.StatusBadGateway)
		}
	}
}

func (r *RepomdResource) DefaultHandler() http.HandlerFunc {
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

func (r *RepomdResource) RPMHandler() http.HandlerFunc {
	return func(rw http.ResponseWriter, req *http.Request) {
		download, err := r.decodeDownload(req.Body)
		if err != nil {
			status := http.StatusBadRequest
			http.Error(rw, http.StatusText(status), status)
			log.Print(err)
			return
		}
		log.Printf("rpm-handler processed for: %s", download.URL)
	}
}

func (r *RepomdResource) MetadataHandler() http.HandlerFunc {
	return func(rw http.ResponseWriter, req *http.Request) {
		download, err := r.decodeDownload(req.Body)
		if err != nil {
			status := http.StatusBadRequest
			http.Error(rw, http.StatusText(status), status)
			log.Print(err)
			return
		}

		dataLocation := download.GetLink("data")

		r.RepomdService.ProcessMetadata(
			"http://vault.centos.org/6.4/os/x86_64",
			download.URL,
			dataLocation)

		log.Printf("metadata-handler processed for: %s", download.URL)

		rw.WriteHeader(http.StatusOK)
	}
}
