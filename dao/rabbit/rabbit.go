package rabbit

import (
	"fmt"

	"github.com/streadway/amqp"
)

// todo: rabbit 的连接需要设计

var (
	conn *amqp.Connection
)

func Init(user, pass, host string, port int) (err error) {
	url := fmt.Sprintf("amqp://%s:%s@%s:%d",
		user, pass, host, port)
	conn, err = amqp.Dial(url)
	return
}
