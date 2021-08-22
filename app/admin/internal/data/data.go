package data

import (
	"github.com/Casper-Mars/cloud-restaurant/app/admin/internal/conf"
	"github.com/elastic/go-elasticsearch/v7"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewEsClient, NewCommentRepo)

// Data .
type Data struct {
	Es *elasticsearch.Client
}

// NewData .
func NewData(c *conf.Data, logger log.Logger,
	es *elasticsearch.Client,
) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")

	}
	return &Data{
		Es: es,
	}, cleanup, nil
}
