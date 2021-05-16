package main

import "github.com/panjf2000/gnet"

type Agent struct {
	nodeID int
	userID int

	conn gnet.Conn

	// 可能的其他附加信息
}
