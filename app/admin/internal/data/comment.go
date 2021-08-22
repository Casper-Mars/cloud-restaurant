package data

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/Casper-Mars/cloud-restaurant/app/admin/internal/biz"
	"github.com/go-kratos/kratos/v2/log"
)

type commentRepo struct {
	log  *log.Helper
	data *Data
}

func NewCommentRepo(data *Data, logger log.Logger) biz.CommentRepo {
	return &commentRepo{
		log:  log.NewHelper(logger),
		data: data,
	}
}

func (c commentRepo) Query(ctx context.Context, do *biz.CommentQueryDO) ([]*biz.CommentDO, error) {
	query := func(cond *biz.CommentQueryDO) map[string]interface{} {
		tmp := make(map[string]interface{})
		if do.UserName != "" {
			tmp["Username"] = do.UserName
		}
		if do.FoodName != "" {
			tmp["FoodName"] = do.FoodName
		}
		if do.Comment != "" {
			tmp["Comment"] = do.Comment
		}
		if len(tmp) != 0 {
			return map[string]interface{}{
				"match": tmp,
			}
		}
		return map[string]interface{}{
			"match_all": tmp,
		}
	}
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(map[string]interface{}{
		"query": query(do),
	}); err != nil {
		return nil, err
	}
	search, err := c.data.Es.Search(
		c.data.Es.Search.WithContext(ctx),
		c.data.Es.Search.WithIndex("comment"),
		c.data.Es.Search.WithSize(int(do.Size)),
		c.data.Es.Search.WithFrom(int(do.Index)-1),
		c.data.Es.Search.WithBody(&buf),
		c.data.Es.Search.WithTrackTotalHits(true),
		c.data.Es.Search.WithPretty(),
	)
	if err != nil {
		return nil, err
	}
	defer search.Body.Close()
	r := make(map[string]interface{})
	if err = json.NewDecoder(search.Body).Decode(&r); err != nil {
		return nil, err
	}
	r2 := r["hits"].(map[string]interface{})["hits"].([]interface{})
	size := len(r2)
	if size == 0 {
		return nil, fmt.Errorf("empty")
	}
	result := make([]*biz.CommentDO, size)
	for i, hit := range r2 {
		m := hit.(map[string]interface{})["_source"].(map[string]interface{})
		result[i] = &biz.CommentDO{
			FoodId:   uint64(m["FoodId"].(float64)),
			FoodName: m["FoodName"].(string),
			UserId:   uint64(m["UserId"].(float64)),
			UserName: m["UserName"].(string),
			Comment:  m["Comment"].(string),
		}
	}
	return result, nil
}
