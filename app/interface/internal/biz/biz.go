package biz

import (
	"context"
	v1 "github.com/Casper-Mars/cloud-restaurant/api/user/v1"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/google/wire"
)

// ProviderSet is biz providers.
var ProviderSet = wire.NewSet(NewAuthUsecase, NewUserClient, NewHealthUsecase)

func NewUserClient() v1.UserClient {
	con, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint("dns:///user.default.svc.cluster.local"),
	)
	if err != nil {
		panic(err)
	}
	return v1.NewUserClient(con)
}
