package main

import (
	"fmt"
	"log"

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
	msg, err := ch.Consume(q.Name,
		"consumerim",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		fmt.Println("warning message cant publish")
	}
	consumeCh := make(chan struct{})
	go func() {
		for d := range msg {
			log.Printf("Received a message: %s", d.Body)
			if string(d.Body) == "bitir" {
				close(consumeCh)
				log.Println("bitiriliyor")
			}
		}
	}()
	<-consumeCh
}
