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
	id, err := c.uc.Add(ctx, &biz.CommentDO{
		UserId:  req.UserId,
		FoodId:  req.FoodId,
		Comment: req.Comment,
	})
	if err != nil {
		return nil, err
	}
	return &v1.CommentModifyResp{
		Id: id,
	}, nil
}

func (c CommentService) ListComment(ctx context.Context, empty *emptypb.Empty) (*v1.CommentList, error) {
	list, err := c.uc.List(ctx)
	if err != nil {
		return nil, err
	}
	result := make([]*v1.CommentList_CommentListItem, len(list))
	for i, k := range list {
		result[i] = &v1.CommentList_CommentListItem{
			Comment: k.Comment,
			UserId:  k.UserId,
			FoodId:  k.FoodId,
		}
	}
	return &v1.CommentList{
		Items: result,
	}, nil
}
