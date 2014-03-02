package rpm

type Location struct {
	Href string `xml:"href,attr"`
}

type Checksum struct {
	Type  string `xml:"type,attr"`
	PkgId string `xml:"pkgid,attr,omitempty"`
	Value string `xml:",chardata"`
}
