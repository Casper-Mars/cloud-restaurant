package service

import (
	"context"
	v1 "github.com/Casper-Mars/cloud-restaurant/api/food/v1"
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
	do := &biz.FoodDO{
		Name: req.Name,
	}
	err := f.uc.AddFood(ctx, do)
	if err != nil {
		return nil, err
	}
	return &v1.FoodModifyResp{
		Id: do.Id,
	}, nil
}

func (f FoodService) ListByIds(ctx context.Context, req *v1.ListFoodByIdReq) (*v1.FoodListResp, error) {

	foods, err := f.uc.ListByIds(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	result := make([]*v1.FoodListResp_FoodItem, len(foods))
	for i, k := range foods {
		result[i] = &v1.FoodListResp_FoodItem{
			Id:   k.Id,
			Name: k.Name,
		}
	}
	return &v1.FoodListResp{
		Items: result,
	}, nil
}
