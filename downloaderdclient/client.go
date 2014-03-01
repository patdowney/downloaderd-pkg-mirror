package downloaderdclient

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	api "github.com/patdowney/downloaderd-deb/downloaderdapi"
)

type Client struct {
	URL string
}

func NewDownloaderdClient(url string) *Client {
	c := &Client{URL: url}

	return c
}

func (c *Client) remoteDownload(downloadRequest *api.Request) (*api.Request, error) {
	jsonBytes, err := json.Marshal(downloadRequest)
	byteReader := bytes.NewReader(jsonBytes)

	res, err := http.Post(c.URL, "application/json", byteReader)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode == http.StatusAccepted || res.StatusCode == http.StatusOK {

		jsonDecoder := json.NewDecoder(res.Body)
		var newRequest api.Request
		err = jsonDecoder.Decode(&newRequest)
		if err != nil {
			panic(err)
		}
		return &newRequest, nil
	}
	return nil, errors.New(fmt.Sprintf("unexpected status code %d returned from %s when requesting %s", res.StatusCode, c.URL, downloadRequest.URL))
}

func (c *Client) RequestDownloadWithCallback(url string, callback string) (*api.Request, error) {
	r := api.NewRequest(url, callback)

	return c.RequestDownload(&r)
}

func (c *Client) RequestDownload(r *api.Request) (*api.Request, error) {
	return c.remoteDownload(r)
}

/*
func (c *Client) GetDownloadReader(download *api.Download) (io.ReadCloser, err) {
	return nil, nil
}
*/
