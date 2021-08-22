package data

import (
	"github.com/Casper-Mars/cloud-restaurant/app/admin/internal/conf"
	"github.com/elastic/go-elasticsearch/v7"
	"net/http"
	"time"
)

func NewEsClient(c *conf.Data) *elasticsearch.Client {
	config := elasticsearch.Config{
		Addresses: []string{c.Elasticsearch.Url},
		Username:  c.Elasticsearch.Username,
		Password:  c.Elasticsearch.Password,
		Transport: &http.Transport{
			ResponseHeaderTimeout: time.Second * 10,
		},
	}
	client, err := elasticsearch.NewClient(config)
	if err != nil {
		panic(err)
	}
	return client
}
