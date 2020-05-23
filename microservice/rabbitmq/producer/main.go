package main

import (
	"fmt"
	"log"
	"time"

	"github.com/streadway/amqp"
)

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	defer conn.Close()
	if err != nil {
		log.Fatal("rabbitmq server connection error")
	}
	ch, err := conn.Channel()
	if err != nil {
		log.Fatal("channel cant created")
	}
	q, err := ch.QueueDeclare(
		"hello", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	if err != nil {
		log.Fatal("que creating error")
	}
	for i := 0; i < 10; i++ {
		err = ch.Publish(
			"",
			q.Name,
			false,
			false,
			amqp.Publishing{
				ContentType: "text/plain",
				Body:        []byte(fmt.Sprintf("test body:%v", i)),
			},
		)
		time.Sleep(time.Second * time.Duration(i))
		if err != nil {
			fmt.Println("warning message cant publish")
		}
	}
	err = ch.Publish(
		"",
		q.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte("bitir"),
		},
	)
	if err != nil {
		fmt.Println("warning message cant publish")
	}
}
