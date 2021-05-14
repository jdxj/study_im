package main

import (
	"log"

	"github.com/jdxj/study_im/proto/head"
)

func (c *Client) handle() {
	for {
		data := <-c.receiveChan

		switch msg := data.(type) {
		case *head.Head:
			c.handleHead(msg)
		}
	}
}

func (c *Client) handleHead(headMsg *head.Head) {
	log.Printf("receive *head.Head:\n%s\n", headMsg)
	c.seq = headMsg.Seq
}
