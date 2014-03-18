package main

import (
	"log"

	"github.com/patdowney/downloaderd-pkg-mirror/deb"
	client "github.com/patdowney/downloaderd-pkg-mirror/downloaderdclient"
	dh "github.com/patdowney/downloaderd-pkg-mirror/http"
	"github.com/patdowney/downloaderd-pkg-mirror/rpm"
)

func webmain(c *client.Client, listenAddress string) {
	s := dh.NewServer(&dh.HTTPConfig{ListenAddress: listenAddress})

	ds := deb.NewDebianService(c, listenAddress)
	dr := dh.NewDebianResource(ds)
	s.AddResource("/deb", dr)

	rs := rpm.NewRepomdService(c, listenAddress)
	rr := dh.NewRepomdResource(rs)
	s.AddResource("/rpm", rr)

	err := s.ListenAndServe()
	log.Print(err)
}
