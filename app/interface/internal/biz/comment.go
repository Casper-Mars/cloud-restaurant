package biz

import (
	"context"
	"encoding/json"
	"errors"
	commentv1 "github.com/Casper-Mars/cloud-restaurant/api/comment/v1"
	foodv1 "github.com/Casper-Mars/cloud-restaurant/api/food/v1"
	userv1 "github.com/Casper-Mars/cloud-restaurant/api/user/v1"
	"github.com/Casper-Mars/cloud-restaurant/app/interface/internal/data"
	"github.com/Casper-Mars/cloud-restaurant/pkg/status"
	"github.com/elastic/go-elasticsearch/v7/esapi"
	"github.com/go-kratos/kratos/v2/log"
	"golang.org/x/sync/errgroup"
	"google.golang.org/protobuf/types/known/emptypb"
	"strconv"
	"strings"
	"time"
)

var commentTopic = "comment"

type CommentDO struct {
	Id       uint64 `json:"Id"`
	UserId   uint64 `json:"UserId"`
	UserName string `json:"UserName"`
	FoodId   uint64 `json:"FoodId"`
	FoodName string `json:"FoodName"`
	Comment  string `json:"Comment"`
}

type CommentUsecase struct {
	log  *log.Helper
	uc   userv1.UserClient
	fc   foodv1.FoodClient
	cc   commentv1.CommentClient
	data *data.Data
}

func NewCommentUsecase(logger log.Logger, uc userv1.UserClient, fc foodv1.FoodClient, cc commentv1.CommentClient, data *data.Data) *CommentUsecase {
	return &CommentUsecase{
		log:  log.NewHelper(logger),
		uc:   uc,
		fc:   fc,
		cc:   cc,
		data: data,
	}
}

func (receiver CommentUsecase) AddComment(ctx context.Context, do CommentDO) (uint64, error) {
	//todo add transactions
	comment, err := receiver.cc.AddComment(ctx, &commentv1.CommentAddReq{
		Comment: do.Comment,
		UserId:  do.UserId,
		FoodId:  do.FoodId,
	})
	if err != nil {
		return 0, err
	}

	commentDO := CommentDO{
		UserId:  do.UserId,
		FoodId:  do.FoodId,
		Comment: do.Comment,
	}

	users, err1 := receiver.uc.ListUserByIds(ctx, &userv1.ListUserByIdReq{Id: []uint64{do.UserId}})
	foods, err2 := receiver.fc.ListByIds(ctx, &foodv1.ListFoodByIdReq{Id: []uint64{do.FoodId}})

	if err1 == nil && err2 == nil {
		commentDO.UserName = users.Items[0].Name
		commentDO.FoodName = foods.Items[0].Name
		marshal, err := json.Marshal(commentDO)
		if err != nil {
			receiver.log.Errorf("Json error:%s", err)
		} else {
			request := esapi.IndexRequest{
				Index:      commentTopic,
				DocumentID: strconv.Itoa(int(comment.Id)),
				Refresh:    "true",
				Body:       strings.NewReader(string(marshal)),
				Timeout:    time.Second,
			}
			response, err := request.Do(context.Background(), receiver.data.Es)
			if err != nil {
				receiver.log.Error(err)
			} else {
				defer response.Body.Close()
			}
		}
	} else {
		receiver.log.Errorf("Can not get user or food:\n%s\n%s", err1, err2)
	}

	//marshal, err := json.Marshal(comment)
	//if err != nil {
	//	return 0, err
	//}
	//err = receiver.data.Kafka.Producer.Produce(&kafka.Message{
	//	TopicPartition: kafka.TopicPartition{Topic: &commentTopic, Partition: kafka.PartitionAny},
	//	Value:          marshal,
	//}, nil)
	//if err != nil {
	//	receiver.log.Error(err)
	//}
	//receiver.data.Kafka.Producer.Flush(1000)
	return comment.Id, nil
}

func (receiver CommentUsecase) ListComment(ctx context.Context) ([]*CommentDO, error) {
	/*search for the comment*/
	comments, err := receiver.cc.ListComment(ctx, &emptypb.Empty{})
	if err != nil {
		if errors.As(err, &status.ErrQueryEmpty) {
			return make([]*CommentDO, 0), nil
		}
		return nil, err
	}
	result := make([]*CommentDO, len(comments.Items))
	userIds := make([]uint64, len(comments.Items))
	foodIds := make([]uint64, len(comments.Items))
	for i, k := range comments.Items {
		result[i] = &CommentDO{
			UserId:  k.UserId,
			FoodId:  k.FoodId,
			Comment: k.Comment,
		}
		userIds[i] = k.UserId
		foodIds[i] = k.FoodId
	}
	userMap := make(map[uint64]string, len(userIds))
	foodMap := make(map[uint64]string, len(userIds))
	/*search the user and food info*/
	timeout, _ := context.WithTimeout(context.Background(), time.Second*5)
	group, _ := errgroup.WithContext(timeout)
	group.Go(func() error {
		users, err2 := receiver.uc.ListUserByIds(ctx, &userv1.ListUserByIdReq{
			Id: userIds,
		})
		if err2 == nil {
			userMap = make(map[uint64]string, len(users.Items))
			for _, u := range users.Items {
				userMap[u.Id] = u.Name
			}
		}
		return nil
	})
	group.Go(func() error {
		foods, err2 := receiver.fc.ListByIds(ctx, &foodv1.ListFoodByIdReq{
			Id: foodIds,
		})
		if err2 == nil {
			foodMap = make(map[uint64]string, len(foods.Items))
			for _, f := range foods.Items {
				foodMap[f.Id] = f.Name
			}
		}
		return nil
	})
	group.Wait()
	/*fill the user and food info */
	for _, k := range result {
		if userMap != nil {
			k.UserName = userMap[k.UserId]
		}
		if foodMap != nil {
			k.FoodName = foodMap[k.FoodId]
		}
	}
	return result, nil
}
