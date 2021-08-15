package biz

import "github.com/go-kratos/kratos/v2/log"

type FoodDO struct {
}

type FoodRepo interface {
	AddFood(foodDO *FoodDO) error
	UpdateFood(foodDO *FoodDO) error
}

type FoodUsecase struct {
	repo FoodRepo
	log  *log.Helper
}

func NewFoodUsecase(repo FoodRepo, logger log.Logger) *FoodUsecase {
	return &FoodUsecase{
		repo: repo,
		log:  log.NewHelper(logger),
	}
}
