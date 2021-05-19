package rabbit

import (
	"fmt"
	"reflect"
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
		"test_rabbit", 5672)
	err := b.Connect()
	if err != nil {
		t.Fatalf("%s\n", err)
	}

	err = b.Subscribe(nil)
	if err != nil {
		t.Fatalf("%s\n", err)
	}

	for {
		time.Sleep(time.Second)

		err = b.Publish("test_rabbit", nil, []byte("hah"))
		if err != nil {
			t.Fatalf("%s\n", err)
		}
	}
}

func TestPubAndSub(t *testing.T) {
	b1 := New("guest", "guest", "127.0.0.1", "b1", 5672)
	err := b1.Connect()
	if err != nil {
		t.Fatalf("%s\n", err)
	}

	b2 := New("guest", "guest", "127.0.0.1", "b2", 5672)
	err = b2.Connect()
	if err != nil {
		t.Fatalf("%s\n", err)
	}

	err = b2.Subscribe(func(m map[string]interface{}, msg []byte) error {
		logicType := reflect.TypeOf(m["logic"])
		fmt.Printf("logicType: %s\n", logicType.Kind())
		fmt.Printf("%s\n", msg)
		return err
	})
	if err != nil {
		t.Fatalf("%s\n", err)
	}

	for {
		time.Sleep(time.Second)
		m := make(map[string]interface{})
		m["logic"] = uint32(2)
		err = b1.Publish("b2", m, []byte("hello"))
		if err != nil {
			t.Fatalf("%s\n", err)
		}
	}
}
