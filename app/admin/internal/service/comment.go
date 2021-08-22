package service

import (
	"context"
	"github.com/Casper-Mars/cloud-restaurant/api/admin/v1"
	"github.com/Casper-Mars/cloud-restaurant/app/admin/internal/biz"
	"github.com/go-kratos/kratos/v2/log"
)

type CommentService struct {
	v1.UnimplementedCommentServer
	log *log.Helper
	uc  *biz.CommentUsecase
}

func NewCommentService(logger log.Logger, uc *biz.CommentUsecase) *CommentService {
	return &CommentService{
		log: log.NewHelper(logger),
		uc:  uc,
	}
}

func (c CommentService) Query(ctx context.Context, req *v1.CommentQueryReq) (*v1.CommentQueryResp, error) {
	query := &biz.CommentQueryDO{
		Comment:  req.Comment.GetValue(),
		UserName: req.UserName.GetValue(),
		FoodName: req.FoodName.GetValue(),
		Index:    1,
		Size:     10,
	}
	if req.Index.GetValue() != 0 {
		query.Index = req.Index.GetValue()
	}
	if req.Size.GetValue() != 0 {
		query.Size = req.Size.GetValue()
	}
	comments, err := c.uc.Query(ctx, query)
	if err != nil {
		c.log.Errorf("Error occur while query comment: %s", err)
		return nil, nil
	}
	result := make([]*v1.CommentItem, len(comments))
	for i, k := range comments {
		result[i] = &v1.CommentItem{
			UserId:   k.UserId,
			UserName: k.UserName,
			FoodId:   k.FoodId,
			FoodName: k.FoodName,
			Comment:  k.Comment,
		}
	}
	return &v1.CommentQueryResp{
		Items: result,
	}, nil
}
