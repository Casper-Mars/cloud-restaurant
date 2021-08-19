package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type User struct {
	Id       uint64
	Name     string
	Phone    string
	Password string
}

type UserRepo interface {
	AddUser(ctx context.Context, user *User) error
	UpdateUser(ctx context.Context, user *User) error
	List(ctx context.Context) ([]User, error)
	ListByIds(ctx context.Context, ids []uint64) ([]User, error)
}

type UserUsecase struct {
	repo UserRepo
	log  *log.Helper
}

func NewUserUsecase(repo UserRepo, logger log.Logger) *UserUsecase {
	return &UserUsecase{
		repo: repo,
		log:  log.NewHelper(logger),
	}
}

func (receiver UserUsecase) AddUser(ctx context.Context, user *User) error {
	return receiver.repo.AddUser(ctx, user)
}

func (receiver UserUsecase) UpdateUser(ctx context.Context, user *User) error {
	return receiver.repo.AddUser(ctx, user)
}

func (receiver UserUsecase) List(ctx context.Context) []User {
	list, err := receiver.repo.List(ctx)
	if err != nil {
		receiver.log.Error(err)
		return nil
	}
	return list
}

func (receiver UserUsecase) ListByIds(ctx context.Context, ids []uint64) []*User {
	users, err := receiver.repo.ListByIds(ctx, ids)
	if err != nil {
		receiver.log.Error(err)
		return nil
	}
	result := make([]*User, len(users))

	for i, k := range users {
		result[i] = &User{
			Id:    k.Id,
			Name:  k.Name,
			Phone: k.Phone,
		}
	}
	return result
}
