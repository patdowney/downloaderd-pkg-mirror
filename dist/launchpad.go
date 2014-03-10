package distribution

import (
	"fmt"
	"net/url"
)

type LaunchpadRelease struct {
	ID string

	Host        string // e.g. http://ppa.launchpad.net
	User        string // e.g. brightbox
	Archive     string // e.g. passenger
	ReleaseName string // e.g. precise, hardy

}

func (r *LaunchpadRelease) BaseURL() (*url.URL, error) {
	return url.Parse(fmt.Sprintf("%s/%s/%s/ubuntu/", r.Host, r.User, r.Archive))
}

func (r *LaunchpadRelease) ReleaseBaseURL() (*url.URL, error) {
	base, err := r.BaseURL()
	if err != nil {
		return nil, err
	}
	return base.Parse(fmt.Sprintf("dists/%s/", r.ReleaseName))
}

func (r *LaunchpadRelease) PackageURL(relativePath string) (*url.URL, error) {
	base, err := r.BaseURL()
	if err != nil {
		return nil, err
	}
	return base.Parse(relativePath)
}

func (r *LaunchpadRelease) ManifestURL() (*url.URL, error) {
	base, err := r.ReleaseBaseURL()
	if err != nil {
		return nil, err
	}
	return base.Parse("Release")
}

func (r *LaunchpadRelease) ManifestSignatureURL() (*url.URL, error) {
	base, err := r.ReleaseBaseURL()
	if err != nil {
		return nil, err
	}
	return base.Parse("Release.gpg")
}
