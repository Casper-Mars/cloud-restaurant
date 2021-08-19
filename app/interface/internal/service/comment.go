package service

import (
	"context"
	"errors"
	v1 "github.com/Casper-Mars/cloud-restaurant/api/interface/v1"
	"github.com/Casper-Mars/cloud-restaurant/app/interface/internal/biz"
	"github.com/go-kratos/kratos/v2/log"
	"google.golang.org/protobuf/types/known/emptypb"
)

type CommentService struct {
	v1.UnimplementedCommentServer
	cc  *biz.CommentUsecase
	log *log.Helper
}

func (c CommentService) AddComment(ctx context.Context, req *v1.CommentAddReq) (*v1.CommentModifyResp, error) {

	comment, err := c.cc.AddComment(ctx, biz.CommentDO{
		Comment: req.Comment.GetValue(),
		UserId:  req.UserId.GetValue(),
		FoodId:  req.FoodId.GetValue(),
	})
	if err != nil {
		return nil, err
	}
	return &v1.CommentModifyResp{
		CommentId: comment,
	}, nil
}

func (c CommentService) ListComment(ctx context.Context, empty *emptypb.Empty) (*v1.ListCommentResp, error) {
	comment := c.cc.ListComment(ctx)
	if comment == nil {
		return nil, errors.New("")
	}
	result := make([]*v1.ListCommentResp_ListCommentItem, len(comment))
	for i, k := range comment {
		result[i] = &v1.ListCommentResp_ListCommentItem{
			UserId:   k.UserId,
			UserName: k.UserName,
			FoodId:   k.FoodId,
			FoodName: k.FoodName,
			Comment:  k.Comment,
		}
	}
	return &v1.ListCommentResp{
		Items: result,
	}, nil
}
