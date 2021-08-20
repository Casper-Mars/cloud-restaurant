package biz

import (
	"context"
	foodv1 "github.com/Casper-Mars/cloud-restaurant/api/food/v1"
	"github.com/go-kratos/kratos/v2/log"
	"google.golang.org/protobuf/types/known/emptypb"
)

type FoodDO struct {
	Id   uint64
	Name string
}

type FoodUsecase struct {
	log *log.Helper
	fc  foodv1.FoodClient
}

func NewFoodUsecase(logger log.Logger, fc foodv1.FoodClient) *FoodUsecase {
	return &FoodUsecase{
		log: log.NewHelper(logger),
		fc:  fc,
	}
}

func (receiver FoodUsecase) AddFood(ctx context.Context, do *FoodDO) error {
	add, err := receiver.fc.Add(ctx, &foodv1.AddFoodReq{
		Name: do.Name,
	})
	if err != nil {
		return err
	}
	do.Id = add.Id
	return nil
}

func (receiver FoodUsecase) List(ctx context.Context) ([]*FoodDO, error) {
	list, err := receiver.fc.List(ctx, &emptypb.Empty{})
	if err != nil {
		return nil, err
	}
	result := make([]*FoodDO, len(list.Items))
	for i, k := range list.Items {
		result[i] = &FoodDO{
			Id:   k.Id,
			Name: k.Name,
		}
	}
	return result, nil
}
