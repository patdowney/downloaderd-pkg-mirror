package distribution

import (
	"fmt"
	"testing"
)

func TestFedoraReleaseManifestURL(t *testing.T) {
	expectedManifestURL := "http://dl.fedoraproject.org/pub/fedora/linux/releases/20/Everything/x86_64/os/repodata/repomd.xml"

	r := FedoraRelease{
		Host:         "http://dl.fedoraproject.org/pub/fedora/linux/",
		Branch:       "releases",
		Release:      "20",
		Tree:         "Everything",
		Architecture: "x86_64",
		Group:        "os"}

	manifestURL, err := r.ManifestURL()
	if err != nil {
		t.Error(err)
	}

	if expectedManifestURL != manifestURL.String() {
		t.Error(fmt.Sprintf("%s != %s", expectedManifestURL, manifestURL.String()))
	}
}

func TestFedoraReleasePackageURL(t *testing.T) {
	expectedPackageURL := "http://dl.fedoraproject.org/pub/fedora/linux/releases/20/Everything/x86_64/os/Packages/0/0xFFFF-0.3.9-10.fc20.x86_64.rpm"
	r := FedoraRelease{
		Host:         "http://dl.fedoraproject.org/pub/fedora/linux/",
		Branch:       "releases",
		Release:      "20",
		Tree:         "Everything",
		Architecture: "x86_64",
		Group:        "os"}

	packageURL, err := r.PackageURL("Packages/0/0xFFFF-0.3.9-10.fc20.x86_64.rpm")
	if err != nil {
		t.Error(err)
	}

	if expectedPackageURL != packageURL.String() {
		t.Error(fmt.Sprintf("%s != %s", expectedPackageURL, packageURL.String()))
	}
}
