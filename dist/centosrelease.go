package distribution

import (
	"fmt"
	"net/url"
)

type CentOSRelease struct {
	ID string

	Host         string // e.g. http://vault.centos.org/
	Architecture string // e.g. x86_64,Source,i386
	Version      string // e.g. 6.4, 6.5
	Group        string // e.g. os, contrib, updates, xen4, fasttrack, cr, centosplus, extras
}

func (r *CentOSRelease) BaseURL() (*url.URL, error) {
	return url.Parse(r.Host)
}

func (r *CentOSRelease) ReleaseBaseURL() (*url.URL, error) {
	base, err := r.BaseURL()
	if err != nil {
		return nil, err
	}
	return base.Parse(fmt.Sprintf("%s/%s/%s/", r.Version, r.Group, r.Architecture))
}

func (r *CentOSRelease) PackageURL(relativePath string) (*url.URL, error) {
	base, err := r.ReleaseBaseURL()
	if err != nil {
		return nil, err
	}
	return base.Parse(relativePath)
}

func (r *CentOSRelease) ManifestURL() (*url.URL, error) {
	base, err := r.ReleaseBaseURL()
	if err != nil {
		return nil, err
	}
	return base.Parse("repodata/repomd.xml")
}
