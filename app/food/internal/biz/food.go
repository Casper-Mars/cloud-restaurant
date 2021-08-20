package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type FoodDO struct {
	Id   uint64
	Name string
}

type FoodRepo interface {
	AddFood(ctx context.Context, foodDO *FoodDO) error
	List(ctx context.Context) ([]*FoodDO, error)
	ListByIds(ctx context.Context, ids []uint64) ([]*FoodDO, error)
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

func (receiver FoodUsecase) AddFood(ctx context.Context, do *FoodDO) error {
	return receiver.repo.AddFood(ctx, do)
}

func (receiver FoodUsecase) List(ctx context.Context) ([]*FoodDO, error) {
	return receiver.repo.List(ctx)
}

func (receiver FoodUsecase) ListByIds(ctx context.Context, ids []uint64) ([]*FoodDO, error) {
	return receiver.repo.ListByIds(ctx, ids)
}
