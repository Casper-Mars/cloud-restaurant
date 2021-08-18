package service

import (
	context "context"
	"github.com/Casper-Mars/cloud-restaurant/api/user/v1"
	"github.com/Casper-Mars/cloud-restaurant/user/internal/biz"
	"github.com/go-kratos/kratos/v2/log"
)

type UserService struct {
	v1.UnimplementedUserServer

	uc  *biz.UserUsecase
	log *log.Helper
}

func NewUserService(uc *biz.UserUsecase, logger log.Logger) *UserService {
	return &UserService{
		uc:  uc,
		log: log.NewHelper(logger),
	}
}

func (u UserService) ListUer(ctx context.Context, req *v1.OnePageUserReq) (*v1.UserListResp, error) {
	panic("implement me")
}

func (u UserService) AddUser(ctx context.Context, req *v1.AddUserReq) (*v1.UserModifyResp, error) {
	user := &biz.User{
		Password: req.Password,
		Name:     req.Name,
		Phone:    req.Phone,
	}
	err := u.uc.AddUser(ctx, user)
	if err != nil {
		return nil, err
	}
	return &v1.UserModifyResp{
		Id: user.Id,
	}, nil
}
