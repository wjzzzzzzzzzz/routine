package main

import (
	"fmt"
	"github.com/streadway/amqp"
	"strconv"
	"time"
)

func main() {
	go produce()
	go consume(1)
	go consume(2)
	select {}
}
func produce() {
	conn, err := amqp.Dial("amqp://wjz:123456@47.101.132.76:5672")
	if err != nil {
		fmt.Print(err)
	}
	fmt.Print(conn)
	ch, err := conn.Channel()
	if err != nil {
		return
	}
	defer ch.Close()
	q, err := ch.QueueDeclare("test number", false, false, false, false, nil)
	for i := 0; i < 60; i++ {
		body := "hello, world " + strconv.Itoa(i)
		err = ch.Publish(
			"",
			q.Name,
			false,
			false,
			amqp.Publishing{
				ContentType: "text/plain",
				Body:        []byte(body),
			})
		time.Sleep(time.Second)
	}
}
func consume(tag uint) {
	conn, err := amqp.Dial("amqp://wjz:123456@47.101.132.76:5672")
	if err != nil {
		fmt.Print(err)
	}
	defer conn.Close()
	ch, err := conn.Channel()
	if err != nil {
		return
	}
	defer ch.Close()
	q, err := ch.QueueDeclare(
		"test number", // name
		false,         // durable //持久化
		false,         // delete when usused
		false,         // exclusive
		false,         // no-wait
		nil,           // arguments
	)
	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack //  消息的时候调用d.Ack(false)来告诉RabbitMQ该消息已经消费
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	for {
		select {
		case m := <-msgs:
			fmt.Println(tag, string(m.Body))
			//m.Ack(false)//手动消息 防止某个消费者失效时，重新发送消息给某个消费者
		}
	}

}
