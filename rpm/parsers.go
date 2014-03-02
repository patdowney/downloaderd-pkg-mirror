package rpm

import (
	"encoding/xml"
	"io"
)

func ParseMetadata(body io.Reader) (*Metadata, error) {
	metadata := Metadata{}

	xmlDecoder := xml.NewDecoder(body)
	err := xmlDecoder.Decode(&metadata)

	if err != nil {
		return nil, err
	}

	return &metadata, nil
}

func ParseRepomd(body io.Reader) (*Repomd, error) {
	repomd := Repomd{}

	xmlDecoder := xml.NewDecoder(body)
	err := xmlDecoder.Decode(&repomd)

	if err != nil {
		return nil, err
	}

	return &repomd, nil

}
