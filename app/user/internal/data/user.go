package data

import (
	context "context"
	"github.com/Casper-Mars/cloud-restaurant/app/user/internal/biz"
	"github.com/go-kratos/kratos/v2/log"
)

type userRepo struct {
	data *Data
	log  *log.Helper
}

func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return userRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (u userRepo) AddUser(ctx context.Context, user *biz.User) error {
	panic("implement me")
}

func (u userRepo) UpdateUser(ctx context.Context, user *biz.User) error {
	panic("implement me")
}

func (u userRepo) List(ctx context.Context) ([]biz.User, error) {
	panic("implement me")
}
