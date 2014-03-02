package main

import (
	"fmt"
	"log"

	"github.com/patdowney/downloaderd-pkg/deb"
	client "github.com/patdowney/downloaderd-pkg/downloaderdclient"
	dh "github.com/patdowney/downloaderd-pkg/http"
	"github.com/patdowney/downloaderd-pkg/rpm"
)

func webmain(c *client.Client, listenPort int, listenHost string) {
	listenAddress := fmt.Sprintf("%s:%d", listenHost, listenPort)
	s := dh.NewServer(&dh.HTTPConfig{ListenAddress: listenAddress})

	ds := deb.NewDebianService(c)
	dr := dh.NewDebianResource(ds)
	s.AddResource("/deb", dr)

	rs := rpm.NewRepomdService(c)
	rr := dh.NewRepomdResource(rs)
	s.AddResource("/rpm", rr)

	err := s.ListenAndServe()
	log.Print(err)
}
