package downloaderdapi

type Request struct {
	URL          string           `json:"url"`
	Checksum     string           `json:"checksum,omitempty"`
	ChecksumType string           `json:"checksum_type,omitempty"`
	DownloadID   string           `json:"download_id,omitempty"`
	Metadata     DownloadMetadata `json:"metadata,omitempty"`
	Errors       []DownloadError  `json:"errors,omitempty"`
	Callback     string           `json:"callback,omitempty"`
	Links        []Link           `json:"links,omitempty"`
}

func NewRequest(url string, callback string) Request {
	return Request{
		URL:      url,
		Callback: callback}
}
