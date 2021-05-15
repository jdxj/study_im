package main

import (
	"net/http"
	_ "net/http/pprof"
	"os"
	"unsafe"
)

func main() {
	switch os.Args[1] {
	case "server":
		go http.ListenAndServe("0.0.0.0:6060", nil)

		StartServer()
	default:
		StartClient()
	}
}

func IsLittleEndian() bool {
	n := 0x1234
	f := *((*byte)(unsafe.Pointer(&n)))
	return (f ^ 0x34) == 0
}
