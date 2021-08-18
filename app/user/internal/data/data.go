package data

import (
	"context"
	"github.com/Casper-Mars/cloud-restaurant/user/internal/conf"
	"github.com/Casper-Mars/cloud-restaurant/user/internal/data/ent"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	_ "github.com/lib/pq"
	"time"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewUserRepo)

// Data .
type Data struct {
	client *ent.Client
}

// NewData .
func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {
	open, err := ent.Open(c.Database.Driver, c.Database.Source)
	if err != nil {
		return nil, nil, err
	}
	ctx, _ := context.WithTimeout(context.Background(), time.Second*2)
	err = open.Schema.Create(ctx)
	if err != nil {
		return nil, nil, err
	}
	cleanup := func() {
		helper := log.NewHelper(logger)
		helper.Info("closing the data resources")
		if err2 := open.Close(); err2 != nil {
			helper.Errorf("Error occur when closing database: %s\n", err2.Error())
		}
	}
	return &Data{
		client: open,
	}, cleanup, nil
}
