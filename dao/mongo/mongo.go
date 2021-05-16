package mongo

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	client *mongo.Client
)

func Init(user, pass, dbname, host string, port int) (err error) {
	uri := fmt.Sprintf("mongodb://%s:%s@%s:%d",
		user, pass, host, port)
	client, err = mongo.Connect(context.Background(), options.Client().ApplyURI(uri))
	if err != nil {
		return
	}
	err = client.Ping(nil, nil)
	return
}
