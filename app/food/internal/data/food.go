package data

import (
	"context"
	"github.com/Casper-Mars/cloud-restaurant/app/food/internal/biz"
	"github.com/Casper-Mars/cloud-restaurant/app/food/internal/data/ent/food"
	"github.com/go-kratos/kratos/v2/log"
)

type foodRepo struct {
	data *Data
	log  *log.Helper
}

func (f foodRepo) AddFood(ctx context.Context, foodDO *biz.FoodDO) error {
	save, err := f.data.client.Food.Create().
		SetName(foodDO.Name).
		Save(ctx)
	if err != nil {
		return err
	}
	foodDO.Id = uint64(save.ID)
	return nil
}

func (f foodRepo) List(ctx context.Context) ([]*biz.FoodDO, error) {
	all, err := f.data.client.Food.Query().All(ctx)
	if err != nil {
		return nil, err
	}
	result := make([]*biz.FoodDO, len(all))
	for i, k := range all {
		result[i] = &biz.FoodDO{
			Id:   uint64(k.ID),
			Name: k.Name,
		}
	}
	return result, nil
}

func (f foodRepo) ListByIds(ctx context.Context, ids []uint64) ([]*biz.FoodDO, error) {
	targets := make([]int, len(ids))
	for i, k := range ids {
		targets[i] = int(k)
	}
	all, err := f.data.client.Food.Query().Where(food.IDIn(targets...)).All(ctx)
	if err != nil {
		return nil, err
	}
	result := make([]*biz.FoodDO, len(all))
	for i, k := range all {
		result[i] = &biz.FoodDO{
			Id:   uint64(k.ID),
			Name: k.Name,
		}
	}
	return result, nil
}

func NewFoodRepo(data *Data, logger log.Logger) biz.FoodRepo {
	return &foodRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}
