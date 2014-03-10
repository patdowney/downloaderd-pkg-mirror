package distribution

import (
	"fmt"
	"testing"
)

func TestUbuntuReleaseManifestURL(t *testing.T) {
	expectedManifestURL := "http://archive.ubuntu.com/ubuntu/dists/precise/Release"

	r := DebianRelease{Host: "http://archive.ubuntu.com/ubuntu/", Name: "precise"}

	manifestURL, err := r.ManifestURL()
	if err != nil {
		t.Error(err)
	}

	if expectedManifestURL != manifestURL.String() {
		t.Error(fmt.Sprintf("%s != %s", expectedManifestURL, manifestURL.String()))
	}
}

func TestUbuntuReleasePackageURL(t *testing.T) {
	expectedPackageURL := "http://archive.ubuntu.com/ubuntu/pool/main/f/firefox/abrowser_20.0+build1-0ubuntu0.12.04.3_amd64.deb"

	r := DebianRelease{Host: "http://archive.ubuntu.com/ubuntu/", Name: "precise"}

	packageURL, err := r.PackageURL("pool/main/f/firefox/abrowser_20.0+build1-0ubuntu0.12.04.3_amd64.deb")
	if err != nil {
		t.Error(err)
	}

	if expectedPackageURL != packageURL.String() {
		t.Error(fmt.Sprintf("%s != %s", expectedPackageURL, packageURL.String()))
	}
}
