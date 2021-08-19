package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type CommentDO struct {
	Comment string
	UserId  uint64
	FoodId  uint64
}

type CommentRepo interface {
	Add(ctx context.Context, do *CommentDO) (uint64, error)
	List(ctx context.Context) ([]*CommentDO, error)
}

type CommentUsecase struct {
	repo CommentRepo
	log  *log.Helper
}

func NewCommentUsecase(repo CommentRepo, logger log.Logger) *CommentUsecase {
	return &CommentUsecase{
		repo: repo,
		log:  log.NewHelper(logger),
	}
}

func (receiver CommentUsecase) Add(ctx context.Context, do *CommentDO) (uint64, error) {
	return receiver.repo.Add(ctx, do)
}

func (receiver CommentUsecase) List(ctx context.Context) ([]*CommentDO, error) {
	return receiver.repo.List(ctx)
}
