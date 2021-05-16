package rabbit

import (
	"fmt"
	"testing"
)

func TestInit(t *testing.T) {
	err := Init("guest", "guest", "127.0.0.1", 5672)
	if err != nil {
		t.Fatalf("%s\n", err)
	}
	fmt.Printf("ok")
	conn.Close()
}
