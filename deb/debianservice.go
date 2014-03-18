package deb

import (
	"compress/gzip"
	"fmt"
	"log"
	"net/http"
	"path/filepath"

	api "github.com/patdowney/downloaderd-pkg-mirror/downloaderdapi"
	client "github.com/patdowney/downloaderd-pkg-mirror/downloaderdclient"
	"github.com/patdowney/godebiancontrol"
)

type DebianService struct {
	Client        *client.Client
	ListenAddress string
}

func NewDebianService(c *client.Client, listenAddress string) *DebianService {
	s := &DebianService{Client: c, ListenAddress: listenAddress}
	return s
}

func (s *DebianService) ProcessPackages(baseURL string, originalURL string, packagesURL string) error {
	res, err := http.Get(packagesURL)
	defer res.Body.Close()
	if err != nil {
		return err
	}
	if res.StatusCode == 200 {
		gzipReader, err := gzip.NewReader(res.Body)
		if err != nil {
			return err
		}

		packagesData := ParsePackagesContent(gzipReader)
		log.Printf("%s: packages:%d", originalURL, len(packagesData))

		for _, p := range packagesData {
			_, err := s.processPackage(baseURL, p)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (s *DebianService) CallbackForFile(filePath string) string {
	baseFilename := filepath.Base(filePath)
	if baseFilename == "Packages.gz" {
		return fmt.Sprintf("http://%s/deb/packages-handler", s.ListenAddress)
	} else if baseFilename == "Release" {
		return fmt.Sprintf("http://%s/deb/release-handler", s.ListenAddress)
	} else if filepath.Ext(baseFilename) == ".deb" {
		return fmt.Sprintf("http://%s/deb/deb-handler", s.ListenAddress)
	} else {
		return fmt.Sprintf("http://%s/deb/default-handler", s.ListenAddress)

	}
}

func (s *DebianService) processPackage(baseURL string, packageData godebiancontrol.Paragraph) (*api.Request, error) {
	packagePath := packageData["Filename"]
	packageDebURL := fmt.Sprintf("%s/%s", baseURL, packagePath)

	checksumType := "SHA256"

	packageDownloadRequest := api.Request{
		URL:          packageDebURL,
		Checksum:     packageData[checksumType],
		ChecksumType: checksumType,
		Callback:     s.CallbackForFile(packagePath)}

	remoteRequest, err := s.Client.RequestDownload(&packageDownloadRequest)

	return remoteRequest, err
}

/*
  and now the release stuff
*/
