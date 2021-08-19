package service

import (
	"context"
	v1 "github.com/Casper-Mars/cloud-restaurant/api/interface/v1"
	"github.com/Casper-Mars/cloud-restaurant/app/interface/internal/biz"
	"google.golang.org/protobuf/types/known/emptypb"
)

type HealthService struct {
	v1.UnimplementedHealthServer
	hc *biz.HealthUsecase
}

func NewHealthService(hc *biz.HealthUsecase) *HealthService {
	return &HealthService{
		hc: hc,
	}
}

func (h HealthService) TestUser(ctx context.Context, empty *emptypb.Empty) (*emptypb.Empty, error) {
	err := h.hc.TestUser()
	if err != nil {
		panic(err)
	}
	return &emptypb.Empty{}, nil
}
