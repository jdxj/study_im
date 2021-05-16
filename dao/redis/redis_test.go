package redis

import (
	"context"
	"testing"
)

func TestInit(t *testing.T) {
	err := Init("", "127.0.0.1", 6379, 0)
	if err != nil {
		t.Fatalf("%s\n", err)
	}
	defer client.Close()

	err = client.Set(context.Background(), "abc", "def", -1).Err()
	if err != nil {
		t.Fatalf("%s\n", err)
	}
}
