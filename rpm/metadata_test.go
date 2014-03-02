package rpm

import (
	"strings"
	"testing"
)

func TestMetadataParsing(t *testing.T) {
	metadata := `
<?xml version="1.0" encoding="UTF-8"?>
<metadata xmlns="http://linux.duke.edu/metadata/common" xmlns:rpm="http://linux.duke.edu/metadata/rpm" packages="6381">
<package type="rpm">
  <name>migrationtools</name>
  <arch>noarch</arch>
  <version epoch="0" ver="47" rel="7.el6"/>
  <checksum type="sha256" pkgid="YES">ec9d688898b0f6438c2d422229071e9cbb26ba2fe4d451fb62e813f7195dc649</checksum>
  <summary>Migration scripts for LDAP</summary>
  <description>The MigrationTools are a set of Perl scripts for migrating users, groups,
aliases, hosts, netgroups, networks, protocols, RPCs, and services from
existing nameservices (flat files, NIS, and NetInfo) to LDAP.</description>
  <packager>CentOS BuildSystem &lt;http://bugs.centos.org&gt;</packager>
  <url>http://www.padl.com/OSS/MigrationTools.html</url>
  <time file="1309667414" build="1282666158"/>
  <size package="25192" installed="106663" archive="111420"/>
<location href="Packages/migrationtools-47-7.el6.noarch.rpm"/>
  <format>
    <rpm:license>BSD</rpm:license>
    <rpm:vendor>CentOS</rpm:vendor>
    <rpm:group>System Environment/Daemons</rpm:group>
    <rpm:buildhost>c6b2.bsys.dev.centos.org</rpm:buildhost>
    <rpm:sourcerpm>migrationtools-47-7.el6.src.rpm</rpm:sourcerpm>
    <rpm:header-range start="1384" end="8748"/>
    <rpm:provides>
      <rpm:entry name="migrationtools" flags="EQ" epoch="0" ver="47" rel="7.el6"/>
    </rpm:provides>
    <rpm:requires>
      <rpm:entry name="/bin/sh"/>
      <rpm:entry name="/usr/bin/perl"/>
      <rpm:entry name="ldif2ldbm"/>
      <rpm:entry name="openldap-clients"/>
      <rpm:entry name="perl"/>
    </rpm:requires>
  </format>
</package>
</metadata>
`

	r := strings.NewReader(metadata)
	m, err := ParseMetadata(r)
	if err != nil {
		t.Error(err)
	}

	packagesLen := len(m.Packages)
	expectedLen := 1
	if packagesLen != expectedLen {
		t.Errorf(`len(v.Packages) = %q want %q`, packagesLen, expectedLen)
	}
}
