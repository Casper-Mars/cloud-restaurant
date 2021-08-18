package biz

import (
	"context"
	v1 "github.com/Casper-Mars/cloud-restaurant/api/user/v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

type HealthUsecase struct {
	uc v1.UserClient
}

func NewHealthUsecase(uc v1.UserClient) *HealthUsecase {
	return &HealthUsecase{
		uc: uc,
	}
}

func (receiver HealthUsecase) TestUser() error {
	_, err := receiver.uc.Heath(context.Background(), &emptypb.Empty{})
	if err != nil {
		return err
	}
	return nil
}
