package service

import (
	"context"
	v1 "github.com/Casper-Mars/cloud-restaurant/api/interface/v1"
	"github.com/Casper-Mars/cloud-restaurant/app/interface/internal/biz"
	"github.com/go-kratos/kratos/v2/log"
	"google.golang.org/protobuf/types/known/emptypb"
)

type FoodService struct {
	v1.UnimplementedFoodServer

	log *log.Helper
	fc  *biz.FoodUsecase
}

func NewFoodService(logger log.Logger, fc *biz.FoodUsecase) *FoodService {
	return &FoodService{
		log: log.NewHelper(logger),
		fc:  fc,
	}
}

func (f FoodService) AddFood(ctx context.Context, req *v1.AddFoodReq) (*v1.FoodModifyResp, error) {
	do := &biz.FoodDO{
		Name: req.Name,
	}
	err := f.fc.AddFood(ctx, do)
	if err != nil {
		return nil, err
	}
	return &v1.FoodModifyResp{
		Id: do.Id,
	}, nil
}

func (f FoodService) List(ctx context.Context, empty *emptypb.Empty) (*v1.FoodList, error) {
	list, err := f.fc.List(ctx)
	if err != nil {
		return nil, err
	}
	result := make([]*v1.FoodList_FoodListItem, len(list))
	for i, k := range list {
		result[i] = &v1.FoodList_FoodListItem{
			Id:   k.Id,
			Name: k.Name,
		}
	}
	return &v1.FoodList{
		Items: result,
	}, nil
}
