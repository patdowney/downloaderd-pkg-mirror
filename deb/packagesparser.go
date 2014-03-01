package deb

import (
	"github.com/patdowney/godebiancontrol"
	"io"
)

type PackageData godebiancontrol.Paragraph
type PackagesData []godebiancontrol.Paragraph

func ParsePackagesContent(body io.Reader) PackagesData {
	paragraphs, _ := ParseControl(body)
	return PackagesData(paragraphs)
}
