package service

import (
	context "context"
	"fmt"
	"github.com/Casper-Mars/cloud-restaurant/api/user/v1"
	"github.com/Casper-Mars/cloud-restaurant/app/user/internal/biz"
	"github.com/go-kratos/kratos/v2/log"
	"google.golang.org/protobuf/types/known/emptypb"
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
	list := u.uc.List(ctx)
	result := make([]*v1.UserListResp_UserListItem, len(list))
	for i, k := range list {
		result[i] = &v1.UserListResp_UserListItem{
			Id:    k.Id,
			Name:  k.Name,
			Phone: k.Phone,
		}
	}
	return &v1.UserListResp{
		Items: result,
	}, nil
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

func (u UserService) Heath(ctx context.Context, in *emptypb.Empty) (*emptypb.Empty, error) {
	fmt.Println("get a call")
	return &emptypb.Empty{}, nil
}

func (u UserService) ListUserByIds(ctx context.Context, in *v1.ListUserByIdReq) (*v1.UserListResp, error) {
	users := u.uc.ListByIds(ctx, in.Id)
	if users == nil {
		return nil, nil
	}
	result := make([]*v1.UserListResp_UserListItem, len(users))
	for i, k := range users {
		result[i] = &v1.UserListResp_UserListItem{
			Id:    k.Id,
			Name:  k.Name,
			Phone: k.Phone,
		}
	}
	return &v1.UserListResp{
		Items: result,
	}, nil
}
