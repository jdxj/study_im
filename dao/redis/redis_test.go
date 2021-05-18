package redis

import (
	"context"
	"fmt"
	"log"
	"reflect"
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

func TestMain(m *testing.M) {
	err := Init("", "127.0.0.1", 6379, 0)
	if err != nil {
		log.Fatalln(err)
	}
	defer client.Close()

	m.Run()
}

func TestHSet(t *testing.T) {
	m := map[string]interface{}{
		"name": "abc",
		"age":  123,
	}
	err := client.HMSet(context.Background(), "test_hset", m).Err()
	if err != nil {
		t.Fatalf("%s\n", err)
	}
}

func TestHGet(t *testing.T) {
	result, err := client.HMGet(context.Background(), "test_hset", "name", "age").Result()
	if err != nil {
		t.Fatalf("%s\n", err)
	}

	typ := reflect.TypeOf(result[1])
	fmt.Println(typ)
}

func TestSession_Set(t *testing.T) {
	s := &Session{
		NodeID:   1,
		ClientID: 23,
		UserID:   45,
	}
	err := s.Set()
	if err != nil {
		t.Fatalf("%s\n", err)
	}
}

func TestSession_Get(t *testing.T) {
	s := &Session{UserID: 45}
	err := s.Get()
	if err != nil {
		t.Fatalf("%s\n", err)
	}
	fmt.Printf("%#v, %d\n", *s, s.UserID)
}
