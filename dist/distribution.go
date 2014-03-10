package distribution

import (
	"net/url"
)

type RemoteRepository interface {
	PackageURL(relativePath string) (*url.URL, error)
	ManifestURL() (*url.URL, error)
}
