package data

import (
	"github.com/Casper-Mars/cloud-restaurant/app/food/internal/biz"
	"github.com/go-kratos/kratos/v2/log"
)

type foodRepo struct {
	data *Data
	log  *log.Helper
}

func NewFoodRepo(data *Data, logger log.Logger) biz.FoodRepo {
	return &foodRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (f foodRepo) AddFood(foodDO *biz.FoodDO) error {
	panic("implement me")
}

func (f foodRepo) UpdateFood(foodDO *biz.FoodDO) error {
	panic("implement me")
}
