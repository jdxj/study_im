package main

import (
	"os"

	"github.com/jdxj/study_im/client"
	"github.com/jdxj/study_im/gate"
)

func main() {
	sign := make(chan bool)

	switch os.Args[1] {
	case "server":
		server := gate.New()
		server.Run(sign)
	default:
		c := client.New()
		c.Run(sign)
	}
}
