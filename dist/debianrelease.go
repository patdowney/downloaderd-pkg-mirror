package distribution

import (
	"fmt"
	"net/url"
)

type DebianRelease struct {
	ID   string
	Host string // e.g. http://archive.ubuntu.org/ubuntu/
	Name string // e.g. precise, precise-updates, precise-backports, precise-proposed, precise-security // implied: RelativeStartingPoint string // e.g. "dists/{Name}/Release

	//dists/{Name}/{Group}/{Architecture}"
	// http://archive.ubuntu.com/ubuntu/dists/precise/Release

}

func (r *DebianRelease) BaseURL() (*url.URL, error) {
	u, err := url.Parse(r.Host)
	return u, err
}

func (r *DebianRelease) ReleaseBaseURL() (*url.URL, error) {
	base, err := r.BaseURL()
	if err != nil {
		return nil, err
	}
	return base.Parse(fmt.Sprintf("dists/%s/", r.Name))
}

func (r *DebianRelease) PackageURL(relativePath string) (*url.URL, error) {
	base, err := r.BaseURL()
	if err != nil {
		return nil, err
	}
	return base.Parse(relativePath)
}

func (r *DebianRelease) ManifestURL() (*url.URL, error) {
	base, err := r.ReleaseBaseURL()
	if err != nil {
		return nil, err
	}
	return base.Parse("Release")
}

func (r *DebianRelease) ManifestSignatureURL() (*url.URL, error) {
	base, err := r.ReleaseBaseURL()
	if err != nil {
		return nil, err
	}
	return base.Parse("Release.gpg")
}
