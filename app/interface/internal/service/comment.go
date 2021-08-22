package service

import (
	"context"
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

func NewCommentService(cc *biz.CommentUsecase, logger log.Logger) *CommentService {
	return &CommentService{
		cc:  cc,
		log: log.NewHelper(logger),
	}
}

func (c CommentService) AddComment(ctx context.Context, req *v1.CommentAddReq) (*v1.CommentModifyResp, error) {
	if req.UserId == 0 || req.FoodId == 0 || req.Comment == "" {
		return nil, v1.ErrorParamMiss("missing param")
	}
	comment, err := c.cc.AddComment(ctx, biz.CommentDO{
		Comment: req.Comment,
		UserId:  req.UserId,
		FoodId:  req.FoodId,
	})
	if err != nil {
		return nil, v1.ErrorGenError("error occur while adding comment:%s", err)
	}
	return &v1.CommentModifyResp{
		CommentId: comment,
	}, nil
}

func (c CommentService) ListComment(ctx context.Context, empty *emptypb.Empty) (*v1.ListCommentResp, error) {
	comment, err := c.cc.ListComment(ctx)
	if err != nil {
		return nil, v1.ErrorGenError("error occur while listing comment: %s", err)
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
