package rabbit

import (
	"fmt"
	"time"

	"github.com/streadway/amqp"
)

// todo: rabbit 的连接需要设计

const (
	exchangeName = "im"
)

var (
	conn *amqp.Connection
)

func Init(user, pass, host string, port int) (err error) {
	url := fmt.Sprintf("amqp://%s:%s@%s:%d",
		user, pass, host, port)
	conn, err = amqp.Dial(url)
	return
}

func New(user, pass, host, bindingKey, queueName string, port int) *Broker {
	b := &Broker{
		user:       user,
		pass:       pass,
		host:       host,
		bindingKey: bindingKey,
		queueName:  queueName,
		port:       port,
	}
	return b
}

type Broker struct {
	user string
	pass string
	host string
	port int

	conn       *amqp.Connection
	channel    *amqp.Channel
	bindingKey string // 订阅用
	queueName  string // 订阅用
}

func (b *Broker) Connect() error {
	var err error
	url := fmt.Sprintf("amqp://%s:%s@%s:%d",
		b.user, b.pass, b.host, b.port)
	b.conn, err = amqp.Dial(url)
	if err != nil {
		return err
	}
	b.channel, err = b.conn.Channel()
	if err != nil {
		return err
	}

	err = b.channel.ExchangeDeclare(
		exchangeName,
		"topic",
		true,
		false,
		false,
		false,
		nil,
	)
	return err
}

func (b *Broker) Publish(routingKey string, body []byte) error {
	msg := amqp.Publishing{
		Headers:         nil,
		ContentType:     "text/plain",
		ContentEncoding: "",
		DeliveryMode:    2,
		Priority:        0,
		CorrelationId:   "",
		ReplyTo:         "",
		Expiration:      "",
		MessageId:       "",
		Timestamp:       time.Time{},
		Type:            "",
		UserId:          "",
		AppId:           "",
		Body:            body,
	}
	err := b.channel.Publish(exchangeName, routingKey, true, false, msg)
	return err
}

func (b *Broker) Subscribe() error {
	queue, err := b.channel.QueueDeclare(
		b.queueName, true, false, false, false, nil)
	if err != nil {
		return err
	}
	err = b.channel.QueueBind(queue.Name, b.bindingKey, exchangeName, false, nil)
	if err != nil {
		return err
	}

	msgChan, err := b.channel.Consume(queue.Name, "", false, false, false, false, nil)
	if err != nil {
		return err
	}

	go func() {
		for {
			msg := <-msgChan
			fmt.Printf("receive: %s\n", msg.Body)
			err := msg.Ack(false)
			if err != nil {
				fmt.Println(err)
			}
		}
	}()
	return nil
}
