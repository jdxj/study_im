package mongo

import (
	"testing"

	"go.mongodb.org/mongo-driver/bson"
)

func TestInit(t *testing.T) {
	err := Init("jdxj", "123456", "test", "127.0.0.1", 27017)
	if err != nil {
		t.Fatalf("%s\n", err)
	}
	defer client.Disconnect(nil)

	c := client.Database("test").Collection("test_c")
	d := bson.D{{
		"name",
		123,
	}}
	_, err = c.InsertOne(nil, d)
	if err != nil {
		t.Fatalf("%s\n", err)
	}
}
