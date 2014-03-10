package distribution

import (
	"fmt"
	"testing"
)

func TestLaunchpadReleaseManifestURL(t *testing.T) {
	expectedManifestURL := "http://ppa.launchpad.net/brightbox/passenger/ubuntu/dists/precise/Release"

	r := LaunchpadRelease{
		Host:        "http://ppa.launchpad.net",
		User:        "brightbox",
		Archive:     "passenger",
		ReleaseName: "precise"}

	manifestURL, err := r.ManifestURL()
	if err != nil {
		t.Error(err)
	}

	if expectedManifestURL != manifestURL.String() {
		t.Error(fmt.Sprintf("%s != %s", expectedManifestURL, manifestURL.String()))
	}
}

func TestLaunchpadPackageURL(t *testing.T) {
	expectedPackageURL := "http://ppa.launchpad.net/brightbox/passenger/ubuntu/pool/main/f/fastthread/libfastthread-ruby_1.0.7-1~hardy1_all.deb"

	r := LaunchpadRelease{
		Host:        "http://ppa.launchpad.net",
		User:        "brightbox",
		Archive:     "passenger",
		ReleaseName: "precise"}

	packageURL, err := r.PackageURL("pool/main/f/fastthread/libfastthread-ruby_1.0.7-1~hardy1_all.deb")
	if err != nil {
		t.Error(err)
	}

	if expectedPackageURL != packageURL.String() {
		t.Error(fmt.Sprintf("%s != %s", expectedPackageURL, packageURL.String()))
	}
}
