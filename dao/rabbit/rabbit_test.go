package rabbit

import (
	"fmt"
	"testing"
	"time"
)

func TestInit(t *testing.T) {
	err := Init("guest", "guest", "127.0.0.1", 5672)
	if err != nil {
		t.Fatalf("%s\n", err)
	}
	fmt.Printf("ok")
	conn.Close()
}

func TestNew(t *testing.T) {
	b := New("guest", "guest", "127.0.0.1",
		"test_rabbit", "abc", 5672)
	err := b.Connect()
	if err != nil {
		t.Fatalf("%s\n", err)
	}

	err = b.Subscribe()
	if err != nil {
		t.Fatalf("%s\n", err)
	}

	for {
		time.Sleep(time.Second)

		err = b.Publish("test_rabbit", []byte("hah"))
		if err != nil {
			t.Fatalf("%s\n", err)
		}
	}
}
