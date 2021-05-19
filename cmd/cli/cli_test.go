package main

import (
	"fmt"
	"testing"
)

func TestEscapeR(t *testing.T) {
	fmt.Printf("abc")
	fmt.Printf("\rdef")
}
