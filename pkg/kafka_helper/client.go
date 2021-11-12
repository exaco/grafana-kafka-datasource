package kafka_helper

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type Options struct {
	BootstrapServers string `json:"bootstrapServers"`
	AutoOffsetReset  string `json:"autoOffsetReset"`
	APIKey           string
}

type KafkaClient struct {
	Consumer         *kafka.Consumer
	BootstrapServers string
	AutoOffsetReset  string
}

type KafkaMessage map[string]float64

func NewKafkaClient(options Options) KafkaClient {
	client := KafkaClient{BootstrapServers: options.BootstrapServers, AutoOffsetReset: options.AutoOffsetReset}
	return client
}

func (client *KafkaClient) ConsumerInitialize() {
	var err error
	client.Consumer, err = kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers":  client.BootstrapServers,
		"group.id":           "grafana-kafka-datasource",
		"enable.auto.commit": "false",
		"auto.offset.reset":  client.AutoOffsetReset,
	})
	if err != nil {
		panic(err)
	}
}

func (client *KafkaClient) TopicAssign(topic string, partition int32, offset int64) {
	var err error
	topic_partition := kafka.TopicPartition{
		Topic:     &topic,
		Partition: partition,
		Offset:    kafka.Offset(offset),
		Metadata:  new(string),
		Error:     err,
	}
	partitions := []kafka.TopicPartition{topic_partition}
	err = client.Consumer.Assign(partitions)

	if err != nil {
		panic(err)
	}
	fmt.Printf("Topic assigned.\n")
}

func (client *KafkaClient) ConsumerPull() (KafkaMessage, kafka.Event) {
	var message KafkaMessage
	ev := client.Consumer.Poll(100)
	if ev == nil {
		return message, ev
	}

	switch e := ev.(type) {
	case *kafka.Message:
		json.Unmarshal([]byte(e.Value), &message)
	case kafka.Error:
		// Errors should generally be considered
		// informational, the client will try to
		// automatically recover.
		// But in this example we choose to terminate
		// the application if all brokers are down.
		fmt.Fprintf(os.Stderr, "%% Error: %v: %v\n", e.Code(), e)
		if e.Code() == kafka.ErrAllBrokersDown {
			panic(e)
		}
	default:
	}
	return message, ev
}

func (client KafkaClient) HealthCheck() error {
	client.ConsumerInitialize()
	topic := ""
	_, err := client.Consumer.GetMetadata(&topic, false, 200)
	if err != nil {
		if err.(kafka.Error).Code() == kafka.ErrTransport {
			return err
		}
	}
	return nil
}
