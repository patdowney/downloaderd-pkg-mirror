package distribution

import (
	"fmt"
	"testing"
)

func TestCentOSReleaseManifestURL(t *testing.T) {
	expectedManifestURL := "http://vault.centos.org/6.4/os/x86_64/repodata/repomd.xml"

	r := CentOSRelease{
		Host:         "http://vault.centos.org/",
		Version:      "6.4",
		Group:        "os",
		Architecture: "x86_64"}

	manifestURL, err := r.ManifestURL()
	if err != nil {
		t.Error(err)
	}

	if expectedManifestURL != manifestURL.String() {
		t.Error(fmt.Sprintf("%s != %s", expectedManifestURL, manifestURL.String()))
	}
}

func TestCentOSReleasePackageURL(t *testing.T) {
	expectedPackageURL := "http://vault.centos.org/6.4/os/x86_64/Packages/389-ds-base-1.2.11.15-11.el6.x86_64.rpm"

	r := CentOSRelease{
		Host:         "http://vault.centos.org/",
		Version:      "6.4",
		Group:        "os",
		Architecture: "x86_64"}

	packageURL, err := r.PackageURL("Packages/389-ds-base-1.2.11.15-11.el6.x86_64.rpm")
	if err != nil {
		t.Error(err)
	}

	if expectedPackageURL != packageURL.String() {
		t.Error(fmt.Sprintf("%s != %s", expectedPackageURL, packageURL.String()))
	}
}
