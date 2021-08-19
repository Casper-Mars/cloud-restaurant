package data

import (
	"context"
	"github.com/Casper-Mars/cloud-restaurant/app/comment/internal/biz"
	"github.com/go-kratos/kratos/v2/log"
)

type commentRepo struct {
	data *Data
	log  *log.Helper
}

func NewCommentRepo(data *Data, logger log.Logger) biz.CommentRepo {
	return commentRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (c commentRepo) Add(ctx context.Context, do *biz.CommentDO) (uint64, error) {
	save, err := c.data.client.Comment.Create().
		SetComment(do.Comment).
		SetUserID(do.UserId).
		SetFoodID(do.FoodId).
		Save(ctx)
	if err != nil {
		return 0, err
	}
	return uint64(save.ID), nil
}

func (c commentRepo) List(ctx context.Context) ([]*biz.CommentDO, error) {
	all, err := c.data.client.Comment.Query().All(ctx)
	if err != nil {
		return nil, err
	}
	result := make([]*biz.CommentDO, len(all))
	for i, k := range all {
		result[i] = &biz.CommentDO{
			Comment: k.Comment,
			UserId:  k.UserID,
			FoodId:  k.FoodID,
		}
	}
	return result, nil
}
