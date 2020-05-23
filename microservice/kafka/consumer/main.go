package main

import (
	"log"
	"os"
	"os/signal"

	"github.com/Shopify/sarama"
)

var (
	brokerList        = []string{"localhost:9092"}
	topic             = "important"
	partition         = 0
	offsetType        = -1
	messageCountStart = 0
)

func main() {
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true
	brokers := brokerList
	master, err := sarama.NewConsumer(brokers, config)
	if err != nil {
		log.Panic(err)
	}
	defer func() {
		if err := master.Close(); err != nil {
			log.Panic(err)
		}
	}()
	consumer, err := master.ConsumePartition(topic, 0, sarama.OffsetOldest)
	if err != nil {
		log.Panic(err)
	}
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)
	doneCh := make(chan struct{})
	go func() {
		for {
			select {
			case err := <-consumer.Errors():
				log.Println(err)
			case msg := <-consumer.Messages():
				messageCountStart++
				log.Println("Received messages", string(msg.Key), string(msg.Value))
			case <-signals:
				log.Println("Interrupt is detected")
				doneCh <- struct{}{}
			}
		}
	}()
	<-doneCh
	log.Println("Processed", messageCountStart, "messages")
}
