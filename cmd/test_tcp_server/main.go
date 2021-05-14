package main

import (
	"os"

	main2 "github.com/jdxj/study_im/cmd/gate"

	"github.com/jdxj/study_im/client"
)

func main() {
	sign := make(chan bool)

	switch os.Args[1] {
	case "server":
		server := main2.NewGate()
		server.Run(sign)
	default:
		c := client.New()
		c.Run(sign)
	}
}
