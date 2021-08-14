package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type User struct {
	id       uint64
	name     string
	phone    string
	password string
}

type UserRepo interface {
	AddUser(ctx context.Context, user *User) error
	UpdateUser(ctx context.Context, user *User) error
	List(ctx context.Context) ([]User, error)
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
