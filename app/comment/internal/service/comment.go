package service

import (
	"context"
	v1 "github.com/Casper-Mars/cloud-restaurant/api/comment/v1"
	"github.com/Casper-Mars/cloud-restaurant/app/comment/internal/biz"
	"github.com/go-kratos/kratos/v2/log"
	"google.golang.org/protobuf/types/known/emptypb"
)

type CommentService struct {
	v1.UnimplementedCommentServer

	uc  *biz.CommentUsecase
	log *log.Helper
}

func NewCommentService(uc *biz.CommentUsecase, logger log.Logger) *CommentService {
	return &CommentService{
		uc:  uc,
		log: log.NewHelper(logger),
	}
}

func (c CommentService) AddComment(ctx context.Context, req *v1.CommentAddReq) (*v1.CommentModifyResp, error) {
	panic("implement me")
}

func (c CommentService) ListComment(ctx context.Context, empty *emptypb.Empty) (*v1.CommentList, error) {
	panic("implement me")
}
