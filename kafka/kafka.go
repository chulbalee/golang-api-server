package kafka

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type Kafka struct {
	Producer *kafka.Producer
	iChan    chan kafka.Event

	topic string
}

func (k *Kafka) Init(clienId string, bootstrapServers []string, topic string) {
	fmt.Println(":::KAFKA CONNECTION SETTING")
	fmt.Println(":::KAFKA bootstrapServers: ", strings.Join(bootstrapServers, ","))
	fmt.Println(":::KAFKA topic: ", topic)
	fmt.Println(":::KAFKA clienId: ", clienId)

	if clienId == "" || bootstrapServers == nil || topic == "" {
		panic("err")
	}

	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": strings.Join(bootstrapServers, ","),
		"client.id":         clienId,
		"acks":              "all"})

	if err != nil {
		fmt.Printf("Failed to create producer: %s\n", err)
		os.Exit(1)
	}
	k.Producer = p

	Channel := make(chan kafka.Event, 10000)
	k.iChan = Channel

	k.topic = topic

	go func() {
		for e := range k.Producer.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					fmt.Printf("Failed to deliver message: %v\n", ev.TopicPartition)
				} else {
					fmt.Printf("Successfully produced record to topic %s partition [%d] @ offset %v\n",
						*ev.TopicPartition.Topic, ev.TopicPartition.Partition, ev.TopicPartition.Offset)
				}
			}
		}
	}()

	fmt.Println("::: KAFKA INIT DONE")
}

func (k *Kafka) Produce(data interface{}) {

	var bytes []byte

	bytes, _ = json.Marshal(data)

	err := k.Producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &k.topic, Partition: kafka.PartitionAny},
		Value:          bytes},
		k.iChan,
	)

	if err != nil {
		fmt.Println("kafka produce failed")
	}
}
