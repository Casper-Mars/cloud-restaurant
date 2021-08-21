package data

import (
	"github.com/Casper-Mars/cloud-restaurant/app/interface/internal/conf"
	"github.com/elastic/go-elasticsearch/v7"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(NewData, NewEsClient, NewKafkaData)

type Data struct {
	es    *elasticsearch.Client
	kafka *KafkaData
}

func NewData(logger log.Logger, esc *elasticsearch.Client, kafkaData *KafkaData) (*Data, func(), error) {
	cleanup := func() {
		helper := log.NewHelper(logger)
		helper.Info("closing the data resources")
	}
	return &Data{
		es:    esc,
		kafka: kafkaData,
	}, cleanup, nil
}

func NewEsClient(c *conf.Data) *elasticsearch.Client {
	config := elasticsearch.Config{
		Addresses: []string{c.Elasticsearch.Url},
		Username:  c.Elasticsearch.Username,
		Password:  c.Elasticsearch.Password,
	}
	client, err := elasticsearch.NewClient(config)
	if err != nil {
		panic(err)
	}
	return client
}
