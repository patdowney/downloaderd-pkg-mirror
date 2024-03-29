package rpm

import (
	"compress/gzip"
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"strings"

	api "github.com/patdowney/downloaderd-pkg-mirror/downloaderdapi"
	client "github.com/patdowney/downloaderd-pkg-mirror/downloaderdclient"
)

type RepomdService struct {
	Client        *client.Client
	ListenAddress string
}

func NewRepomdService(c *client.Client, listenAddress string) *RepomdService {
	s := &RepomdService{Client: c, ListenAddress: listenAddress}
	return s
}

func (s *RepomdService) ProcessMetadata(baseURL string, originalURL string, packagesURL string) error {
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

		packagesData, err := ParseMetadata(gzipReader)
		if err != nil {
			return err
		}
		log.Printf("%s: packages:%d", originalURL, len(packagesData.Packages))

		for _, p := range packagesData.Packages {
			_, err := s.processPackage(baseURL, p)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (s *RepomdService) CallbackForFile(filePath string) string {
	baseFilename := filepath.Base(filePath)
	if strings.HasSuffix(baseFilename, "-primary.xml.gz") {
		return fmt.Sprintf("http://%s/rpm/metadata-handler", s.ListenAddress)
	} else if baseFilename == "repomd.xml" {
		return fmt.Sprintf("http://%s/rpm/repomd-handler", s.ListenAddress)
	} else if filepath.Ext(baseFilename) == ".rpm" {
		return fmt.Sprintf("http://%s/rpm/rpm-handler", s.ListenAddress)
	} else {
		return fmt.Sprintf("http://%s/rpm/default-handler", s.ListenAddress)
	}
}

func (s *RepomdService) processPackage(baseURL string, pkg Package) (*api.Request, error) {
	packagePath := pkg.Location.Href
	packageDebURL := fmt.Sprintf("%s/%s", baseURL, packagePath)

	checksumType := "sha256"

	packageDownloadRequest := api.Request{
		URL:          packageDebURL,
		Checksum:     pkg.Checksum.Value,
		ChecksumType: checksumType,
		Callback:     s.CallbackForFile(packagePath)}

	remoteRequest, err := s.Client.RequestDownload(&packageDownloadRequest)

	return remoteRequest, err
}

/*
  and now the release stuff
*/
