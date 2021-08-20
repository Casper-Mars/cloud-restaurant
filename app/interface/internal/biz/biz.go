package biz

import (
	"context"
	commentv1 "github.com/Casper-Mars/cloud-restaurant/api/comment/v1"
	foodv1 "github.com/Casper-Mars/cloud-restaurant/api/food/v1"
	userv1 "github.com/Casper-Mars/cloud-restaurant/api/user/v1"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/google/wire"
)

// ProviderSet is biz providers.
var ProviderSet = wire.NewSet(NewAuthUsecase, NewUserClient, NewFoodClient, NewCommentClient, NewHealthUsecase, NewCommentUsecase, NewFoodUsecase, NewUserUsecase)

func NewUserClient() userv1.UserClient {
	con, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint("dns:///127.0.0.1:9001"),
	)
	if err != nil {
		panic(err)
	}
	return userv1.NewUserClient(con)
}

func NewFoodClient() foodv1.FoodClient {
	con, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint("dns:///127.0.0.1:9002"),
	)
	if err != nil {
		panic(err)
	}
	return foodv1.NewFoodClient(con)
}
func NewCommentClient() commentv1.CommentClient {
	con, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint("dns:///127.0.0.1:9003"),
	)
	if err != nil {
		panic(err)
	}
	return commentv1.NewCommentClient(con)
}
