package main

import (
	"unsafe"

	"github.com/name5566/leaf"
)

func main() {
	//switch os.Args[1] {
	//case "server":
	//	StartServer()
	//default:
	//	StartClient()
	//}
	StartServer()
	leaf.Run()
}

func IsLittleEndian() bool {
	n := 0x1234
	f := *((*byte)(unsafe.Pointer(&n)))
	return (f ^ 0x34) == 0
}
