package rpm

import (
	"encoding/xml"
)

type Metadata struct {
	XMLName      xml.Name  `xml:"metadata"`
	PackageCount uint64    `xml:"packages,attr"`
	Packages     []Package `xml:"package"`
}

type Version struct {
	Epoch   uint64 `xml:"epoch,attr,omitempty"`
	Version string `xml:"ver,attr,omitempty"`
	Release string `xml:"rel,attr,omitempty"`
}

type Time struct {
	File  uint64 `xml:"file,attr"`
	Build uint64 `xml:"build,attr"`
}

type Size struct {
	Package   uint64 `xml:"package,attr"`
	Installed uint64 `xml:"installed,attr"`
	Archive   uint64 `xml:"archive,attr"`
}

type HeaderRange struct {
	Start uint64 `xml:"start,attr"`
	End   uint64 `xml:"end,attr"`
}

type Entry struct {
	XMLName xml.Name `xml:"http://linux.duke.edu/metadata/rpm entry"`
	Version
	Name  string `xml:"name,attr"`
	Flags string `xml:"flags,attr,omitempty"`
}

type Format struct {
	XMLName     xml.Name    `xml:"format"`
	License     string      `xml:"http://linux.duke.edu/metadata/rpm license"`
	Vendor      string      `xml:"http:/linux.duke.edu/metadata/rpm vendor"`
	Group       string      `xml:"http://linux.duke.edu/metadata/rpm group"`
	BuildHost   string      `xml:"http://linux.duke.edu/metadata/rpm buildhost"`
	SourceRPM   string      `xml:"http://linux.duke.edu/metadata/rpm sourcerpm"`
	HeaderRange HeaderRange `xml:"http://linux.duke.edu/metadata/rpm header-range"`
	Provides    []Entry     `xml:"http://linux.duke.edu/metadata/rpm provides>entry"`
	Requires    []Entry     `xml:"http://linux.duke.edu/metadata/rpm requires>entry"`
	Files       []string    `xml:"file"`
}

type Package struct {
	XMLName     xml.Name `xml:"package"`
	Type        string   `xml:"type"`
	Name        string   `xml:"name"`
	Arch        string   `xml:"arch"`
	Version     Version  `xml:"version"`
	Checksum    Checksum `xml:"checksum"`
	Summary     string   `xml:"summary"`
	Description string   `xml:"description"`
	Packager    string   `xml:"packager"`
	Url         string   `xml:"url"`
	Time        Time     `xml:"time"`
	Size        Size     `xml:"size"`
	Location    Location `xml:"location"`
	Format      Format   `xml:"format"`
}
