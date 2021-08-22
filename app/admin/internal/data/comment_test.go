package data

import (
	"context"
	"fmt"
	"github.com/Casper-Mars/cloud-restaurant/app/admin/internal/biz"
	conf2 "github.com/Casper-Mars/cloud-restaurant/app/admin/internal/conf"
	"github.com/go-kratos/kratos/v2/log"
	"os"
	"testing"
)

func TestCommentRepo_Query(t *testing.T) {

	conf := &conf2.Data{
		Elasticsearch: &conf2.Data_Elasticsearch{
			Url:      "http://es-cn-2r42bb1g70041hfoy.public.elasticsearch.aliyuncs.com:9200",
			Username: "elastic",
			Password: "Abcd123456@",
		},
	}
	logger := log.With(log.NewStdLogger(os.Stdout),
		"ts", log.DefaultTimestamp,
		"caller", log.DefaultCaller,
		"service.id", 1,
		"service.name", "1",
		"service.version", "Version",
		"trace_id", log.TraceID(),
		"span_id", log.SpanID(),
	)

	esClient := NewEsClient(conf)
	data, f, err := NewData(conf, logger, esClient)
	if err != nil {
		panic(err)
	}
	defer f()
	repo := NewCommentRepo(data, logger)
	query, err := repo.Query(context.Background(), &biz.CommentQueryDO{
		Index: 0,
		Size:  10,
	})
	if err != nil {
		panic(err)
	}
	for _, k := range query {
		fmt.Printf("%#v\n", *k)
	}
}
