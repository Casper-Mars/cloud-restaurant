package biz

import "github.com/go-kratos/kratos/v2/log"

type CommentDO struct {
}

type CommentRepo interface {
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
