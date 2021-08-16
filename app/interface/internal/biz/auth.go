package biz

import (
	"context"
	v1 "github.com/Casper-Mars/cloud-restaurant/api/user/v1"
	"github.com/go-kratos/kratos/v2/log"
)

type AuthDO struct {
	Phone string
	Pwd   string
}

type UserInfoDO struct {
	Phone string
	Name  string
}

type AuthUsecase struct {
	us  v1.UserClient
	log *log.Helper
}

func NewAuthUsecase(logger log.Logger, us v1.UserClient) *AuthUsecase {
	return &AuthUsecase{
		log: log.NewHelper(logger),
		us:  us,
	}
}

func (receiver AuthUsecase) Registry(ctx context.Context, authInfo *AuthDO) error {
	_, err := receiver.us.AddUser(ctx, &v1.AddUserReq{
		Phone:    authInfo.Phone,
		Password: authInfo.Pwd,
	})
	if err != nil {
		return err
	}
	return nil
}

func (receiver AuthUsecase) Login(ctx context.Context, account, pwd string) (string, error) {
	return "", nil
}

func (receiver AuthUsecase) Logout(ctx context.Context, account string) error {
	return nil
}

func (receiver AuthUsecase) SelfInfo(ctx context.Context, account string) (*UserInfoDO, error) {
	return nil, nil
}
