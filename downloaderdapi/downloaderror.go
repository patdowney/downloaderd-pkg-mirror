package downloaderdapi

import (
	"time"
)

type DownloadError struct {
	Time  time.Time `json:"time"`
	Error string    `json:"error"`
}
