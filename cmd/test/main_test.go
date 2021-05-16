package main

import (
	"encoding/binary"
	"fmt"
	"testing"
	"unsafe"
)

func TestByteOrder(t *testing.T) {
	num := 0x12345678
	first := (*byte)(unsafe.Pointer(&num))
	fmt.Printf("%x\n", *first)
}

func TestByteOrderAPI(t *testing.T) {
	data := []byte{1, 2, 3, 4}
	dataNum := binary.LittleEndian.Uint32(data)
	fmt.Printf("dataNum@: %x\n", dataNum)

	buf2 := [4]byte{1, 2, 3, 4}
	num2 := *(*uint32)(unsafe.Pointer(&buf2))
	fmt.Printf("%x\n", num2)

	buf3 := []byte{1, 2, 3, 4}
	dataNum = binary.BigEndian.Uint32(buf3)
	fmt.Printf("dataNum-: %x\n", dataNum)
	buf4 := make([]byte, 4)
	binary.BigEndian.PutUint32(buf4, dataNum)
	fmt.Printf("buf4: %v\n", buf4)

	num3 := 0x01020304
	array := (*[4]byte)(unsafe.Pointer(&num3))
	fmt.Printf("%v\n", array)

	fmt.Printf("%x\n", binary.LittleEndian.Uint32(array[:]))
	fmt.Printf("%x\n", binary.BigEndian.Uint32(array[:]))

	num4 := 1
	fmt.Printf("1: %v\n", (*[4]byte)(unsafe.Pointer(&num4)))

	num4 = num4 << 8
	fmt.Printf("8: %v\n", (*[4]byte)(unsafe.Pointer(&num4)))

	num4 = num4 << 8
	fmt.Printf("16: %v\n", (*[4]byte)(unsafe.Pointer(&num4)))

	num4 = num4 << 8
	fmt.Printf("24: %v\n", (*[4]byte)(unsafe.Pointer(&num4)))
}

func TestByteOrderAPI2(t *testing.T) {
	var num uint32 = 0x01020304
	buf := make([]byte, 4)
	binary.LittleEndian.PutUint32(buf, num)
	fmt.Printf("%v\n", buf)

	binary.BigEndian.PutUint32(buf, num)
	fmt.Printf("%v\n", buf)
}

type A struct {
	Name string
}

func (a *A) String() string {
	return "abc"
}

func TestAString(t *testing.T) {
	var a *A
	a.String()
}
