package service

import (
	"context"
	"github.com/Casper-Mars/cloud-restaurant/api/food/v1"
	"github.com/Casper-Mars/cloud-restaurant/app/food/internal/biz"
	"github.com/go-kratos/kratos/v2/log"
)

type FoodService struct {
	v1.UnimplementedFoodServer

	uc  *biz.FoodUsecase
	log *log.Helper
}

func NewFoodService(uc *biz.FoodUsecase, logger log.Logger) *FoodService {
	return &FoodService{
		uc:  uc,
		log: log.NewHelper(logger),
	}
}

func (f FoodService) Add(ctx context.Context, req *v1.AddFoodReq) (*v1.FoodModifyResp, error) {
	panic("implement me")
}

func (f FoodService) Update(ctx context.Context, req *v1.UpdateFoodReq) (*v1.FoodModifyResp, error) {
	panic("implement me")
}

func (f FoodService) Page(ctx context.Context, req *v1.OnePageFoodListReq) (*v1.FoodListResp, error) {
	panic("implement me")
}
