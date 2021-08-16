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
	login, err := a.uc.Login(ctx, req.Phone.String(), req.Password.String())
	if err != nil {
		return nil, err
	}
	return &v1.LoginResp{
		Token: login,
	}, nil
}

func (a AuthService) Logout(ctx context.Context, empty *emptypb.Empty) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}

func (a AuthService) SelfInfo(ctx context.Context, empty *emptypb.Empty) (*v1.SelfInfoResp, error) {
	info, err := a.uc.SelfInfo(ctx, "")
	if err != nil {
		return nil, err
	}
	return &v1.SelfInfoResp{
		Phone: info.Phone,
		Name:  info.Name,
	}, nil
}

func (a AuthService) Registry(ctx context.Context, req *v1.RegistryReq) (*emptypb.Empty, error) {
	err := a.uc.Registry(ctx, &biz.AuthDO{
		Phone: req.Phone.String(),
		Pwd:   req.Password.String(),
	})
	return &emptypb.Empty{}, err
}
