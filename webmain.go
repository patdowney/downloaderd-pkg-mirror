package main

import (
	"fmt"
	"log"

	"github.com/patdowney/downloaderd-deb/deb"
	client "github.com/patdowney/downloaderd-deb/downloaderdclient"
	dh "github.com/patdowney/downloaderd-deb/http"
)

func webmain(c *client.Client, listenPort int, listenHost string) {
	listenAddress := fmt.Sprintf("%s:%d", listenHost, listenPort)
	s := dh.NewServer(&dh.HTTPConfig{ListenAddress: listenAddress})

	ds := deb.NewDebianService(c)
	r := dh.NewDebianResource(ds)
	s.AddResource("/deb", r)

	err := s.ListenAndServe()
	log.Print(err)
}
