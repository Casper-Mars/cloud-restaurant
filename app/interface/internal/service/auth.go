package service

import (
	"context"
	"github.com/Casper-Mars/cloud-restaurant/api/interface/v1"
	"github.com/Casper-Mars/cloud-restaurant/app/interface/internal/biz"
	"github.com/go-kratos/kratos/v2/log"
	"google.golang.org/protobuf/types/known/emptypb"
)

type AuthService struct {
	v1.UnimplementedAuthServer

	uc  *biz.AuthUsecase
	log *log.Helper
}

func NewAuthService(uc *biz.AuthUsecase, logger log.Logger) *AuthService {
	return &AuthService{
		uc:  uc,
		log: log.NewHelper(logger),
	}
}

func (a AuthService) Login(ctx context.Context, req *v1.LoginReq) (*v1.LoginResp, error) {
	panic("implement me")
}

func (a AuthService) Logout(ctx context.Context, empty *emptypb.Empty) (*emptypb.Empty, error) {
	panic("implement me")
}

func (a AuthService) SelfInfo(ctx context.Context, empty *emptypb.Empty) (*v1.SelfInfoResp, error) {
	panic("implement me")
}
