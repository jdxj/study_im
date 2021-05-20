package mysql

import (
	"fmt"
	"log"
	"os"
	"testing"
	"time"
)

func TestInit(t *testing.T) {
	err := Init("root", "123456", "test", "127.0.0.1", 3306)
	if err != nil {
		t.Fatalf("%s\n", err)
	}
	defer db.Close()

	rows, err := db.Queryx("select id from test")
	if err != nil {
		t.Fatalf("%s\n", err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		err := rows.Scan(&id)
		if err != nil {
			t.Fatalf("%s\n", err)
		}
		fmt.Printf("%d\n", id)
	}
}

func TestMain(m *testing.M) {
	err := Init("root", "123456", "im", "127.0.0.1", 3306)
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()

	code := m.Run()
	os.Exit(code)
}

func TestUser_Create(t *testing.T) {
	u := &User{
		ID:       0,
		Nickname: "abc",
		Password: "def",
	}
	err := u.Create()
	if err != nil {
		t.Fatalf("%s\n", err)
	}
}

func TestUser_Get(t *testing.T) {
	u := &User{ID: 1}
	exists, err := u.Get()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%#v, exists: %t\n", u, exists)

	u = &User{ID: 0}
	exists, err = u.Get()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%#v, exists: %t\n", u, exists)
}

func TestMessageSend_Insert(t *testing.T) {
	ms := &MessageSend{
		ID:       0,
		FromID:   123,
		ToID:     456,
		Seq:      2,
		Content:  []byte("hello"),
		SendTime: time.Now(),
		SendType: 1,
	}
	err := ms.Insert()
	if err != nil {
		t.Fatalf("%s\n", err)
	}
}

func TestMessageReceive_Insert(t *testing.T) {
	mr := &MessageReceive{
		ID:        0,
		FromID:    123,
		ToID:      456,
		MessageID: 2,
		Flag:      1,
	}
	err := mr.Insert()
	if err != nil {
		t.Fatalf("%s\n", err)
	}
}

func TestMessageReceive_SetRead(t *testing.T) {
	mr := &MessageReceive{
		ToID:      456,
		MessageID: 1,
	}
	err := mr.SetRead()
	if err != nil {
		t.Fatalf("%s\n", err)
	}
}
