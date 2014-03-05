package main

import (
	"flag"
	"fmt"
	"log"
	"net/url"
	"os"

	client "github.com/patdowney/downloaderd-pkg-mirror/downloaderdclient"
)

func ConfigureLogging() {
	log.SetOutput(os.Stderr)
	log.SetFlags(log.LstdFlags | log.Lmicroseconds | log.Lshortfile)
}

func main() {
	var listenPort *int = flag.Int("port", 8081, "port to listen on")
	var listenHost *string = flag.String("host", "", "address to listen on")

	flag.Parse()
	baseUrlArg := flag.Arg(0)
	distName := flag.Arg(1)

	ConfigureLogging()

	downloaderUrl := "http://localhost:8080/request/"
	//--type debian
	//--distBase "http://archive.ubuntu.com/ubuntu/dists/precise-updates"

	var waitForeverChannel chan int

	c := client.NewDownloaderdClient(downloaderUrl)

	go webmain(c, *listenPort, *listenHost)

	baseUrl, _ := url.Parse(baseUrlArg)
	baseDistUrl, _ := baseUrl.Parse("dists/")
	distUrl, _ := baseDistUrl.Parse(fmt.Sprintf("%s/", distName))

	//baseUrl := "http://ftp.us.debian.org/debian/dists/wheezy-updates"
	//baseUrl := "http://archive.ubuntu.com/ubuntu/dists/precise-updates"
	releaseUrl, _ := distUrl.Parse("Release")

	//releaseUrl := fmt.Sprintf("%s/%s", distUrl, "Release")
	//releaseSig := fmt.Sprintf("%s/%s", baseUrl, "Release.gpg")

	_, err := c.RequestDownloadWithCallback(releaseUrl.String(), "http://localhost:8081/deb/release-handler")
	if err != nil {
		panic(err)
	}

	<-waitForeverChannel
}
