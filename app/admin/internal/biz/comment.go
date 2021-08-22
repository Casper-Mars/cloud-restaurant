package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type CommentDO struct {
	UserId   uint64
	UserName string
	FoodId   uint64
	FoodName string
	Comment  string
}

type CommentQueryDO struct {
	Index    uint32
	Size     uint32
	UserName string
	FoodName string
	Comment  string
}

type CommentRepo interface {
	Query(ctx context.Context, do *CommentQueryDO) ([]*CommentDO, error)
}

type CommentUsecase struct {
	log  *log.Helper
	repo CommentRepo
}

func NewCommentUsecase(logger log.Logger, repo CommentRepo) *CommentUsecase {
	return &CommentUsecase{
		log:  log.NewHelper(logger),
		repo: repo,
	}
}

func (receiver CommentUsecase) Query(ctx context.Context, query *CommentQueryDO) ([]*CommentDO, error) {
	comments, err := receiver.repo.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	return comments, nil
}
