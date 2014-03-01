package deb

import (
	"io"

	"github.com/patdowney/godebiancontrol"
)

func ParseControl(body io.Reader) ([]godebiancontrol.Paragraph, error) {
	return godebiancontrol.Parse(body)
}
