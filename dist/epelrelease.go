package distribution

import (
	"fmt"
	"net/url"
)

type EPELRelease struct {
	ID string

	Host         string // e.g. http://dl.fedoraproject.org/pub/fedora/linux/
	Version      string // e.g. 5, 6, 6Server, beta
	Architecture string // e.g. x86_64, SRPMS, i386, armhfp
}

func (r *EPELRelease) BaseURL() (*url.URL, error) {
	return url.Parse(r.Host)
}

func (r *EPELRelease) ReleaseBaseURL() (*url.URL, error) {
	base, err := r.BaseURL()
	if err != nil {
		return nil, err
	}
	return base.Parse(fmt.Sprintf("%s/%s/", r.Version, r.Architecture))
}

func (r *EPELRelease) PackageURL(relativePath string) (*url.URL, error) {
	base, err := r.ReleaseBaseURL()
	if err != nil {
		return nil, err
	}
	return base.Parse(relativePath)
}

func (r *EPELRelease) ManifestURL() (*url.URL, error) {
	base, err := r.ReleaseBaseURL()
	if err != nil {
		return nil, err
	}
	return base.Parse("repodata/repomd.xml")
}
