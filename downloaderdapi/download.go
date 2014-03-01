package downloaderdapi

import (
	"time"
)

type Download struct {
	ID            string    `json:"id"`
	URL           string    `json:"url"`
	TimeRequested time.Time `json:"time_requested"`
	TimeStarted   time.Time `json:"time_started"`
	TimeUpdated   time.Time `json:"time_updated"`
	Size          uint64    `json:"bytes_read"`
	Checksum      string    `json:"checksum"`
	ChecksumType  string    `json:"checksum_type"`
	Finished      bool      `json:"finished"`
	Links         []Link    `json:"links"`
}

func (p *Download) GetLink(relation string) string {
	for _, l := range p.Links {
		if l.Relation == relation {
			return l.Href
		}
	}
	return ""
}
