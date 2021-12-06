package Database

import (
	"bytes"
	"errors"
	"github.com/streadway/amqp"
	"strings"
)

var conn *amqp.Connection
var channel *amqp.Channel
var topics string
var nodes string
var hasMQ bool = false

type Reader interface {
	Read(msg *string) (err error)
}

// 初始化 参数格式：amqp://用户名:密码@地址:端口号/host
func SetupRMQ(rmqAddr string) (err error) {
	if channel == nil {
		conn, err = amqp.Dial(rmqAddr)
		if err != nil {
			return err
		}
		channel, err = conn.Channel()
		if err != nil {
			return err
		}
		hasMQ = true
	}
	return nil
}

// 是否已经初始化
func HasMQ() bool {
	return hasMQ
}

// 测试连接是否正常
func Ping() (err error) {
	if !hasMQ || channel == nil {
		return errors.New("RabbitMQ is not initialize")
	}
	err = channel.ExchangeDeclare("ping.ping", "topic", false, true, false, true, nil)
	if err != nil {
		return err
	}
	msgContent := "ping.ping"
	err = channel.Publish("ping.ping", "ping.ping", false, false, amqp.Publishing{
		ContentType: "text/plain",
		Body:        []byte(msgContent),
	})
	if err != nil {
		return err
	}
	err = channel.ExchangeDelete("ping.ping", false, false)
	return err
}

// 发布消息
func Publish(topic, msg string) (err error) {
	if topics == "" || !strings.Contains(topics, topic) {
		err = channel.ExchangeDeclare(topic, "topic", true, false, false, true, nil)
		if err != nil {
			return err
		}
		topics += "  " + topic + "  "
	}
	err = channel.Publish(topic, topic, false, false, amqp.Publishing{
		ContentType: "text/plain",
		Body:        []byte(msg),
	})
	return nil
}

// 监听接收到的消息
func Receive(topic, node string, reader func(msg *string)) (err error) {
	if topics == "" || !strings.Contains(topics, topic) {
		err = channel.ExchangeDeclare(topic, "topic", true, false, false, true, nil)
		if err != nil {
			return err
		}
		topics += "  " + topic + "  "
	}
	if nodes == "" || !strings.Contains(nodes, node) {
		_, err = channel.QueueDeclare(node, true, false, false, true, nil)
		if err != nil {
			return err
		}
		err = channel.QueueBind(node, topic, topic, true, nil)
		if err != nil {
			return err
		}
		nodes += "  " + node + "  "
	}

	msgs, err := channel.Consume(node, "", true, false, false, false, nil)
	if err != nil {
		return err
	}
	go func() {
		//fmt.Println(*msgs)
		for d := range msgs {
			s := bytesToString(&(d.Body))
			reader(s)
		}
	}()
	return nil
}

// 关闭连接
func Close() {
	channel.Close()
	conn.Close()
	hasMQ = false
}

func bytesToString(b *[]byte) *string {
	s := bytes.NewBuffer(*b)
	r := s.String()
	return &r
}
