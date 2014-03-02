package rpm

import (
	"encoding/xml"
)

type Data struct {
	Type         string   `xml:"type,attr"`
	Location     Location `xml:"location"`
	Checksum     Checksum `xml:"checksum"`
	OpenChecksum Checksum `xml:"open-checksum"`
	Timestamp    float64  `xml:"timestamp"`
	Size         uint64   `xml:"size"`
}

type RepoMd struct {
	XMLName  xml.Name `xml:"repomd"`
	Revision string   `xml:"revision"`
	Data     []Data   `xml:"data"`
}
