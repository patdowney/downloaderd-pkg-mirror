package distribution

import (
	"fmt"
	"testing"
)

func TestEPELReleaseManifestURL(t *testing.T) {
	expectedManifestURL := "http://dl.fedoraproject.org/pub/epel/6Server/x86_64/repodata/repomd.xml"

	r := EPELRelease{
		Host:         "http://dl.fedoraproject.org/pub/epel/",
		Version:      "6Server",
		Architecture: "x86_64"}

	manifestURL, err := r.ManifestURL()
	if err != nil {
		t.Error(err)
	}

	if expectedManifestURL != manifestURL.String() {
		t.Error(fmt.Sprintf("%s != %s", expectedManifestURL, manifestURL.String()))
	}
}

func TestEPELReleasePackageURL(t *testing.T) {
	expectedPackageURL := "http://dl.fedoraproject.org/pub/epel/6Server/x86_64/2ping-2.0-2.el6.noarch.rpm"

	r := EPELRelease{
		Host:         "http://dl.fedoraproject.org/pub/epel/",
		Version:      "6Server",
		Architecture: "x86_64"}

	packageURL, err := r.PackageURL("2ping-2.0-2.el6.noarch.rpm")
	if err != nil {
		t.Error(err)
	}

	if expectedPackageURL != packageURL.String() {
		t.Error(fmt.Sprintf("%s != %s", expectedPackageURL, packageURL.String()))
	}
}
