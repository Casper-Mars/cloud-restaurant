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
	newUser, err := u.data.client.User.Create().
		SetName(user.Name).
		SetPhone(user.Phone).
		SetPwd(user.Password).
		Save(ctx)
	if err != nil {
		return err
	}
	user.Id = newUser.ID
	return nil
}

func (u userRepo) UpdateUser(ctx context.Context, user *biz.User) error {
	_, err := u.data.client.User.UpdateOneID(user.Id).
		SetPhone(user.Phone).
		SetName(user.Name).
		SetPwd(user.Password).
		Save(ctx)
	return err
}

func (u userRepo) List(ctx context.Context) ([]biz.User, error) {
	all, err := u.data.client.User.Query().All(ctx)
	if err != nil {
		return nil, err
	}
	result := make([]biz.User, len(all))
	for i, item := range all {
		result[i] = biz.User{
			Id:       item.ID,
			Name:     item.Name,
			Phone:    item.Phone,
			Password: item.Pwd,
		}
	}
	return result, nil
}
