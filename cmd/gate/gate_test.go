package main

import (
	"fmt"
	"testing"
)

func TestInterface(t *testing.T) {
	var i interface{} = "abc"
	b, _ := i.(int)
	fmt.Printf("%d\n", b)
}
