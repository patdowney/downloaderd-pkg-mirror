package main

import (
	"flag"
	"log"
	"os"

	client "github.com/patdowney/downloaderd-pkg-mirror/downloaderdclient"
)

func ConfigureLogging() {
	log.SetOutput(os.Stderr)
	log.SetFlags(log.LstdFlags | log.Lmicroseconds | log.Lshortfile)
}

func main() {
	var listenAddress *string = flag.String("http", "localhost:8081", "address to listen on")
	var downloaderdUrl *string = flag.String("downloaderd", "http://localhost:8082/request/", "downloaderd request endpoint")

	flag.Parse()

	ConfigureLogging()

	//--type debian
	//--distBase "http://archive.ubuntu.com/ubuntu/dists/precise-updates"

	var waitForeverChannel chan int

	c := client.NewDownloaderdClient(*downloaderdUrl)

	go webmain(c, *listenAddress)

	//		releaseHandler := fmt.Sprintf("http://%s/deb/release-handler", listenAddress)
	<-waitForeverChannel
}
