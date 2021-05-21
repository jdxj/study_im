package rabbit

import (
	"fmt"
	"time"

	"github.com/jdxj/study_im/logger"

	"github.com/streadway/amqp"
)

type Handler func(map[string]interface{}, []byte) error

const (
	exchangeName = "im"
	exchangeKind = "topic"
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

func New(user, pass, host, bindingKey string, port int) *Broker {
	b := &Broker{
		user:       user,
		pass:       pass,
		host:       host,
		bindingKey: bindingKey,
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
	bindingKey string // 订阅用, 不订阅的话不用填写
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
		exchangeKind,
		true,
		false,
		false,
		false,
		nil,
	)
	return err
}

// MIME list: https://www.iana.org/assignments/media-types/media-types.xhtml

func (b *Broker) Publish(routingKey string, headers map[string]interface{}, body []byte) error {
	msg := amqp.Publishing{
		Headers:         headers,
		ContentType:     "application/octet-stream",
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
	// todo: mandatory 如何配置?
	err := b.channel.Publish(exchangeName, routingKey, false, false, msg)
	return err
}

func (b *Broker) Subscribe(h Handler) error {
	// 规定 queue 名称, 避免重启后创建多个 queue
	queueName := fmt.Sprintf("queue.%s", b.bindingKey)
	// autoDelete 设置为 false, 避免队列中仍有消息时而遭到删除
	queue, err := b.channel.QueueDeclare(
		queueName, true, false, false, false, nil)
	if err != nil {
		return err
	}

	err = b.channel.QueueBind(queue.Name, b.bindingKey, exchangeName, false, nil)
	if err != nil {
		return err
	}

	msgChan, err := b.channel.Consume(
		queue.Name, "", false, false, false, false, nil)
	if err != nil {
		return err
	}

	go func() {
		for {
			// todo: 是否使用多个 goroutine?
			msg := <-msgChan
			// todo: 使用 err 来判断是否需要 ack 的做法是否合适?
			err := h(msg.Headers, msg.Body)
			if err != nil {
				logger.Errorf("rabbit handler: %s", err)
				continue
			}

			err = msg.Ack(false)
			if err != nil {
				logger.Errorf("rabbit ack: %s", err)
			}
		}
	}()
	return nil
}
