package distribution

import (
	"fmt"
	"testing"
)

func TestDebianReleaseManifestURL(t *testing.T) {
	expectedManifestURL := "http://ftp.us.debian.org/debian/dists/wheezy-updates/Release"

	r := DebianRelease{Host: "http://ftp.us.debian.org/debian/", Name: "wheezy-updates"}

	manifestURL, err := r.ManifestURL()
	if err != nil {
		t.Error(err)
	}

	if expectedManifestURL != manifestURL.String() {
		t.Error(fmt.Sprintf("%s != %s", expectedManifestURL, manifestURL.String()))
	}
}

func TestDebianReleasePackageURL(t *testing.T) {
	expectedPackageURL := "http://ftp.us.debian.org/debian/pool/main/c/certificatepatrol/xul-ext-certificatepatrol_2.0.14-3+deb7u1_all.deb"
	r := DebianRelease{Host: "http://ftp.us.debian.org/debian/", Name: "wheezy-updates"}

	packageURL, err := r.PackageURL("pool/main/c/certificatepatrol/xul-ext-certificatepatrol_2.0.14-3+deb7u1_all.deb")
	if err != nil {
		t.Error(err)
	}

	if expectedPackageURL != packageURL.String() {
		t.Error(fmt.Sprintf("%s != %s", expectedPackageURL, packageURL.String()))
	}
}
