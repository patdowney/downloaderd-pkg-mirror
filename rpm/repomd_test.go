package rpm

import (
	"strings"
	"testing"
)

func TestRepomdParsing(t *testing.T) {

	repomdData := `
<?xml version="1.0" encoding="UTF-8"?>
<repomd xmlns="http://linux.duke.edu/metadata/repo" xmlns:rpm="http://linux.duke.edu/metadata/rpm">
 <revision>1362488530</revision>
<data type="group">
  <checksum type="sha256">2727fcb43fbe4c1a3588992af8c19e4d97167aee2f6088959221fc285cab6f72</checksum>
  <location href="repodata/2727fcb43fbe4c1a3588992af8c19e4d97167aee2f6088959221fc285cab6f72-c6-x86_64-comps.xml"/>
  <timestamp>1362488701.77</timestamp>
  <size>1189260</size>
</data>
<data type="filelists">
  <checksum type="sha256">3ebb49a547c9d6d96b88ca0befda04c40a083515f33bdfd450640f854f3c87ae</checksum>
  <open-checksum type="sha256">519bc11c1ad67cc2af47ecc61d1c3ff80f775b09f955068c436bd8fd61d71007</open-checksum>
  <location href="repodata/3ebb49a547c9d6d96b88ca0befda04c40a083515f33bdfd450640f854f3c87ae-filelists.xml.gz"/>
  <timestamp>1362488681</timestamp>
  <size>5499724</size>
  <open-size>71465356</open-size>
</data>
</repomd>
`

	r := strings.NewReader(repomdData)
	repomd, err := ParseRepoMd(r)
	if err != nil {
		t.Error(err)
	}

	dataLen := len(repomd.Data)
	expectedLen := 2
	if dataLen != expectedLen {
		t.Errorf(`len(v.Data) = %q want %q`, dataLen, expectedLen)
	}
}
