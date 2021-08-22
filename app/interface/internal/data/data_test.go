package data

import (
	"context"
	"fmt"
	"github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esapi"
	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
	"net/http"
	"strings"
	"testing"
	"time"
)

func TestEs(t *testing.T) {
	config := elasticsearch.Config{
		Addresses: []string{"http://es-cn-2r42bb1g70041hfoy.public.elasticsearch.aliyuncs.com:9200"},
		Username:  "elastic",
		Password:  "Abcd123456@",
		Transport: &http.Transport{
			ResponseHeaderTimeout: time.Second * 10,
		},
	}
	client, err := elasticsearch.NewClient(config)
	if err != nil {
		panic(err)
	}
	info, err := client.Info()
	if err != nil {
		panic(err)
	}

	request := esapi.IndexRequest{
		Index:      "comment",
		DocumentID: "2",
		Timeout:    time.Second,
		Body:       strings.NewReader("{\n          \"FoodId\" : 1,\n          \"FoodName\" : \"宫保鸡丁\",\n          \"UserId\" : 1,\n          \"UserName\" : \"小明\",\n          \"Comment\" : \"great\"\n        }"),
		Refresh:    "true",
	}
	do, err := request.Do(context.Background(), client)
	if err != nil {
		panic(err)
	}
	defer do.Body.Close()

	fmt.Println(info)

}

func TestKafka(t *testing.T) {
	topic := "comment"
	conf := &kafka.ConfigMap{
		"bootstrap.servers":             "120.79.212.109:9093,47.106.100.246:9093",
		"api.version.request":           "true",
		"message.max.bytes":             1000000,
		"linger.ms":                     500,
		"sticky.partitioning.linger.ms": 1000,
		"retries":                       3,
		"retry.backoff.ms":              1000,
		"acks":                          "1",
		"security.protocol":             "sasl_ssl",
		"sasl.username":                 "casperkafka",
		"sasl.password":                 "Abcd123456",
		"sasl.mechanism":                "PLAIN",
		"ssl.ca.location":               "./ca.pem",
	}
	p, err := kafka.NewProducer(conf)
	if err != nil {
		panic(err)
	}
	defer p.Close()

	msg := &kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          []byte("hello world"),
	}

	p.Produce(msg, nil)
	time.Sleep(time.Duration(1) * time.Millisecond)
	p.Flush(15 * 1000)
}
