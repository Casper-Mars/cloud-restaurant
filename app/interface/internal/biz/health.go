package biz

import v1 "github.com/Casper-Mars/cloud-restaurant/api/user/v1"

type HealthUsecase struct {
	uc v1.UserClient
}

func NewHealthUsecase(uc v1.UserClient) *HealthUsecase {
	return &HealthUsecase{
		uc: uc,
	}
}
