package service

import (
	"context"
	"github.com/Casper-Mars/cloud-restaurant/api/food/v1"
)

type FoodService struct {
	v1.UnimplementedFoodServer
}

func NewFoodService() *FoodService {
	return &FoodService{}
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
