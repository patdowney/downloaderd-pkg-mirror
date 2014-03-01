package deb

import (
	"io"
	"regexp"
	"strconv"
	"strings"

	"github.com/patdowney/godebiancontrol"
)

type FileReference struct {
	FileName string
	Size     uint64
	Checksum string
}

type ReleaseData godebiancontrol.Paragraph

func (r ReleaseData) FileReferences(propertyName string) []*FileReference {
	propertyValue := r[propertyName]

	fileRefs := make([]*FileReference, 0)
	if propertyValue != "" {
		fileRefLines := strings.Split(strings.TrimSpace(r[propertyName]), "\n")
		fileRefs = make([]*FileReference, len(fileRefLines))

		for i, line := range fileRefLines {
			if line != "" {
				fileRef := parseFileReference(line)
				fileRefs[i] = fileRef
			}
		}
	}

	return fileRefs
}

func GetFileReferences(paragraph godebiancontrol.Paragraph, propertyName string) []*FileReference {
	fileRefLines := strings.Split(strings.TrimSpace(paragraph[propertyName]), "\n")
	fileRefs := make([]*FileReference, len(fileRefLines))

	for i, line := range fileRefLines {
		if line != "" {
			fileRef := parseFileReference(line)
			fileRefs[i] = fileRef
		}
	}

	return fileRefs
}

func parseFileReference(fileRefLine string) *FileReference {
	fileRefRe := regexp.MustCompile(`\s*([a-z|A-Z|0-9]+?)\s+(\d+)\s+(.+)`)
	match := fileRefRe.FindAllStringSubmatch(fileRefLine, 3)

	checksum, sizeAsString, filename := match[0][1], match[0][2], match[0][3]

	size, _ := strconv.Atoi(sizeAsString)

	return &FileReference{FileName: filename, Checksum: checksum, Size: uint64(size)}
}

func ParseRelease(body io.Reader) ReleaseData {
	paragraphs, _ := ParseControl(body)
	return ReleaseData(paragraphs[0])
}
