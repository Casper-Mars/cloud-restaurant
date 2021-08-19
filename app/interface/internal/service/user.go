package service

import (
	"context"
	v1 "github.com/Casper-Mars/cloud-restaurant/api/interface/v1"
	"github.com/Casper-Mars/cloud-restaurant/app/interface/internal/biz"
	"google.golang.org/protobuf/types/known/emptypb"
)

type UserService struct {
	v1.UnimplementedUserServer

	uc *biz.UserUsecase
}

func NewUserService(uc *biz.UserUsecase) *UserService {
	return &UserService{
		uc: uc,
	}
}

func (u UserService) List(ctx context.Context, empty *emptypb.Empty) (*v1.UserList, error) {
	list, err := u.uc.List(ctx)
	if err != nil {
		return nil, err
	}
	result := make([]*v1.UserList_UserListItem, len(list))
	for i, k := range list {
		result[i] = &v1.UserList_UserListItem{
			Id:    k.Id,
			Name:  k.Name,
			Phone: k.Phone,
		}
	}
	return &v1.UserList{
		Items: result,
	}, nil
}

func (u UserService) Add(ctx context.Context, req *v1.UserAddReq) (*v1.UserModifyResp, error) {
	panic("implement me")
}
