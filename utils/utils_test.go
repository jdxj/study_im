package utils

import (
	"fmt"
	"reflect"
	"testing"
)

type A struct {
	Name string
	Age  int
}

func TestKind(t *testing.T) {
	a := &A{Name: "abc"}
	field := reflect.ValueOf(a).Elem().Field(0)
	if field.Kind() != field.Type().Kind() {
		t.Fatalf("not equal")
	}
}

func TestStruct2Map(t *testing.T) {
	a := &A{
		Name: "abc",
		Age:  123,
	}
	m, err := Struct2Map(a)
	if err != nil {
		t.Fatalf("%s\n", err)
	}
	fmt.Printf("%v\n", m)
}

func TestSetField(t *testing.T) {
	a := &A{}

	m := map[string]interface{}{
		"Name": "abc",
		"Age":  "123",
	}
	for name, value := range m {
		err := SetField(a, name, value)
		if err != nil {
			t.Fatalf("%s\n", err)
		}
	}
	fmt.Printf("%#v\n", a)
}
