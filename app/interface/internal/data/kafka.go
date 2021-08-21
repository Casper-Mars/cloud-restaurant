package data

import (
	"github.com/Casper-Mars/cloud-restaurant/app/interface/internal/conf"
	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
)

type KafkaData struct {
	producer *kafka.Producer
	consumer *kafka.Consumer
}

func NewKafkaData(c *conf.Data) *KafkaData {

	conf := &kafka.ConfigMap{
		"bootstrap.servers":             c.Kafka.Address,
		"api.version.request":           "true",
		"message.max.bytes":             1000000,
		"linger.ms":                     500,
		"sticky.partitioning.linger.ms": 1000,
		"retries":                       3,
		"retry.backoff.ms":              1000,
		"acks":                          "1",
		"security.protocol":             "sasl_ssl",
		"sasl.username":                 c.Kafka.Username,
		"sasl.password":                 c.Kafka.Password,
		"sasl.mechanism":                "PLAIN",
		"ssl.ca.location":               c.Kafka.CaFile,
	}
	var producer *kafka.Producer
	var consumer *kafka.Consumer
	var err error
	if c.Kafka.EnableProducer {
		if producer, err = kafka.NewProducer(conf); err != nil {
			panic(err)
		}
	}
	if c.Kafka.EnableConsumer {
		if consumer, err = kafka.NewConsumer(conf); err != nil {
			panic(err)
		}
	}
	return &KafkaData{
		producer: producer,
		consumer: consumer,
	}
}
