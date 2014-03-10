package distribution

import (
	"fmt"
	"net/url"
)

type FedoraRelease struct {
	ID string

	Host         string // e.g. http://dl.fedoraproject.org/pub/fedora/linux/
	Branch       string // e.g. updates, releases, development
	Architecture string // e.g. x86_64, SRPMS, i386, armhfp
	Release      string // e.g. 20, rawhide
	Tree         string // e.g. Fedora, Everything
	Group        string // e.g. os, debug
}

func (r *FedoraRelease) BaseURL() (*url.URL, error) {
	return url.Parse(r.Host)
}

func (r *FedoraRelease) ReleaseBaseURL() (*url.URL, error) {
	base, err := r.BaseURL()
	if err != nil {
		return nil, err
	}
	return base.Parse(
		fmt.Sprintf("%s/%s/%s/%s/%s/",
			r.Branch,
			r.Release,
			r.Tree,
			r.Architecture,
			r.Group))
}

func (r *FedoraRelease) PackageURL(relativePath string) (*url.URL, error) {
	base, err := r.ReleaseBaseURL()
	if err != nil {
		return nil, err
	}
	return base.Parse(relativePath)
}

func (r *FedoraRelease) ManifestURL() (*url.URL, error) {
	base, err := r.ReleaseBaseURL()
	if err != nil {
		return nil, err
	}
	return base.Parse("repodata/repomd.xml")
}
