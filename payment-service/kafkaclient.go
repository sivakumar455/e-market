package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"golang.org/x/exp/rand"
)

const (
	topic   = "purchases"
	groupID = "kafka-go-123"
)

func getProducerConfig() *kafka.ConfigMap {
	return &kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092",
		//"sasl.username":     "<CLUSTER API KEY>",
		//"sasl.password":     "<CLUSTER API SECRET>",
		//"security.protocol": "SASL_SSL",
		//"sasl.mechanisms":   "PLAIN",
		"acks": "all",
	}
}

func createProducer() (*kafka.Producer, error) {
	producerCfg := getProducerConfig()
	return kafka.NewProducer(producerCfg)
}

func getConsumerConfig() *kafka.ConfigMap {
	return &kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092",
		//"sasl.username":     "<CLUSTER API KEY>",
		//"sasl.password":     "<CLUSTER API SECRET>",
		//"security.protocol": "SASL_SSL",
		//"sasl.mechanisms":   "PLAIN",
		"group.id":          groupID,
		"auto.offset.reset": "earliest",
	}
}

func createConsumer() (*kafka.Consumer, error) {
	consumerCfg := getConsumerConfig()
	return kafka.NewConsumer(consumerCfg)
}

func producer() {

	p, err := createProducer()
	if err != nil {
		fmt.Printf("Failed to create producer: %s", err)
		os.Exit(1)
	}

	users := [...]string{"eabara", "jsmith", "sgarcia", "jbernard", "htanaka", "awalther"}
	items := [...]string{"book", "alarm clock", "t-shirts", "gift card", "batteries"}

	topic := "purchases"

	for n := 0; n < 5; n++ {
		key := users[rand.Intn(len(users))]
		data := items[rand.Intn(len(items))]
		headerMap := map[string]interface{}{"123": []byte("123")}

		publish(p, topic, []byte(key), []byte(data), headerMap)
	}

	// Wait for all messages to be delivered
	p.Flush(15 * 1000)
	p.Close()

}

func publish(p *kafka.Producer, topic string, key []byte, data []byte, headerMap map[string]interface{}) {
	// Go-routine to handle message delivery reports and
	// possibly other event types (errors, stats, etc)
	go handleEvents(p)

	headers := []kafka.Header{}
	for hKey, hVal := range headerMap {
		if value, ok := hVal.([]byte); ok {
			header := kafka.Header{Key: hKey, Value: value}
			headers = append(headers, header)
		} else {
			fmt.Printf("Invalid header for %s is not []byte", hKey)
		}
	}
	topicPartition := kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny}
	msg := &kafka.Message{
		TopicPartition: topicPartition,
		Key:            key,
		Value:          data,
		Headers:        headers,
	}

	p.Produce(msg, nil)

}

func handleEvents(p *kafka.Producer) {
	for e := range p.Events() {
		switch ev := e.(type) {
		case *kafka.Message:
			if ev.TopicPartition.Error != nil {
				fmt.Printf("Failed to deliver message: %v\n", ev.TopicPartition)
			} else {
				fmt.Printf("Produced event to topic %s: key = %-10s value = %s\n",
					*ev.TopicPartition.Topic, string(ev.Key), string(ev.Value))
			}
		}
	}
}

func consumer() {

	c, err := createConsumer()
	if err != nil {
		fmt.Printf("Failed to create consumer: %s\n", err)
		os.Exit(1)
	}

	if err := c.SubscribeTopics([]string{topic}, nil); err != nil {
		fmt.Printf("Failed to subscribe to topics: %s\n", err)
		os.Exit(1)
	}

	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)

	done := make(chan struct{})

	consumeMessages(c, done)

	sig := <-sigchan
	fmt.Printf("Caught signal %v: terminating\n", sig)
	done <- struct{}{}

	if err := c.Close(); err != nil {
		fmt.Printf("Failed to close consumer: %s\n", err)
	}

}

func consumeMessages(c *kafka.Consumer, done chan struct{}) {
	go func() {
		for {
			select {
			case <-done:
				fmt.Println("Terminating...")
				return
			default:
				ev, err := c.ReadMessage(100 * time.Millisecond)
				if err != nil {
					// Errors are informational and automatically handled by the consumer
					continue
				}
				fmt.Printf("Consumed event from topic %s: key = %-10s value = %s\n",
					*ev.TopicPartition.Topic, string(ev.Key), string(ev.Value))
			}
		}
	}()
}
