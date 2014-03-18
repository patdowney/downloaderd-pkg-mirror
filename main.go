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
	var listenAddress *string = flag.String("host", "localhost:8081", "address to listen on")
	var downloaderdUrl *string = flag.String("downloaderd", "http://localhost:8082/request/", "address to listen on")

	flag.Parse()
	baseUrlArg := flag.Arg(0)
	distName := flag.Arg(1)

	ConfigureLogging()

	//--type debian
	//--distBase "http://archive.ubuntu.com/ubuntu/dists/precise-updates"

	var waitForeverChannel chan int

	c := client.NewDownloaderdClient(*downloaderdUrl)

	go webmain(c, *listenAddress)

	baseUrl, _ := url.Parse(baseUrlArg)
	baseDistUrl, _ := baseUrl.Parse("dists/")
	distUrl, _ := baseDistUrl.Parse(fmt.Sprintf("%s/", distName))

	//baseUrl := "http://ftp.us.debian.org/debian/dists/wheezy-updates"
	//baseUrl := "http://archive.ubuntu.com/ubuntu/dists/precise-updates"
	releaseUrl, _ := distUrl.Parse("Release")

	//releaseUrl := fmt.Sprintf("%s/%s", distUrl, "Release")
	//releaseSig := fmt.Sprintf("%s/%s", baseUrl, "Release.gpg")

	releaseHandler := fmt.Sprintf("http://%s/deb/release-handler", listenAddress)
	_, err := c.RequestDownloadWithCallback(releaseUrl.String(), releaseHandler)
	if err != nil {
		panic(err)
	}

	<-waitForeverChannel
}
