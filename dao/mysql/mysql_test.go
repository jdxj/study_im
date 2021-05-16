package mysql

import (
	"fmt"
	"testing"
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
