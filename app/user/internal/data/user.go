package data

import (
	context "context"
	"errors"
	"github.com/Casper-Mars/cloud-restaurant/app/user/internal/biz"
	"github.com/Casper-Mars/cloud-restaurant/app/user/internal/data/ent/user"
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
	user.Id = uint64(newUser.ID)
	return nil
}

func (u userRepo) UpdateUser(ctx context.Context, user *biz.User) error {
	_, err := u.data.client.User.UpdateOneID(int(user.Id)).
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
			Id:       uint64(item.ID),
			Name:     item.Name,
			Phone:    item.Phone,
			Password: item.Pwd,
		}
	}
	return result, nil
}

func (receiver userRepo) ListByIds(ctx context.Context, ids []uint64) ([]biz.User, error) {
	if ids == nil || len(ids) == 0 {
		return nil, errors.New("")
	}
	targets := make([]int, len(ids))
	for i, k := range ids {
		targets[i] = int(k)
	}
	all, err := receiver.data.client.User.Query().Where(user.IDIn(targets...)).All(ctx)
	if err != nil {
		return nil, err
	}
	result := make([]biz.User, len(all))
	for i, u := range all {
		result[i] = biz.User{
			Id:       uint64(u.ID),
			Name:     u.Name,
			Phone:    u.Phone,
			Password: u.Pwd,
		}
	}
	return result, nil
}
