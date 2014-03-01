package downloaderdapi

type DownloadMetadata struct {
	MimeType   string `json:"mime_type"`
	StatusCode int    `json:"http_status_code"`
	Size       uint64 `json:"size"`
}
