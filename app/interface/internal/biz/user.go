package biz

import (
	"context"
	userv1 "github.com/Casper-Mars/cloud-restaurant/api/user/v1"
)

type UserDO struct {
	Id    uint64
	Name  string
	Phone string
}

type UserUsecase struct {
	uc userv1.UserClient
}

func NewUserUsecase(uc userv1.UserClient) *UserUsecase {
	return &UserUsecase{
		uc: uc,
	}
}

func (receiver UserUsecase) List(ctx context.Context) ([]*UserDO, error) {
	uer, err := receiver.uc.ListUer(ctx, &userv1.OnePageUserReq{})
	if err != nil {
		return nil, err
	}
	result := make([]*UserDO, len(uer.Items))
	for i, k := range uer.Items {
		result[i] = &UserDO{
			Id:    k.Id,
			Name:  k.Name,
			Phone: k.Phone,
		}
	}
	return result, nil
}
