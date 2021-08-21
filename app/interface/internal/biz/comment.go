package biz

import (
	"context"
	"encoding/json"
	commentv1 "github.com/Casper-Mars/cloud-restaurant/api/comment/v1"
	foodv1 "github.com/Casper-Mars/cloud-restaurant/api/food/v1"
	userv1 "github.com/Casper-Mars/cloud-restaurant/api/user/v1"
	"github.com/Casper-Mars/cloud-restaurant/app/interface/internal/data"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/go-kratos/kratos/v2/log"
	"google.golang.org/protobuf/types/known/emptypb"
	"sync"
)

var commentTopic = "comment"

type CommentDO struct {
	UserId   uint64
	UserName string
	FoodId   uint64
	FoodName string
	Comment  string
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
	marshal, err := json.Marshal(comment)
	if err != nil {
		return 0, err
	}
	err = receiver.data.Kafka.Producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &commentTopic, Partition: kafka.PartitionAny},
		Value:          marshal,
	}, nil)
	if err != nil {
		receiver.log.Error(err)
	}
	receiver.data.Kafka.Producer.Flush(1000)
	return comment.Id, nil
}

func (receiver CommentUsecase) ListComment(ctx context.Context) []*CommentDO {
	/*search for the comment*/
	comments, err := receiver.cc.ListComment(ctx, &emptypb.Empty{})
	if err != nil {
		receiver.log.Error(err)
		return nil
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
	group := sync.WaitGroup{}
	group.Add(2)
	go func() {
		users, err2 := receiver.uc.ListUserByIds(ctx, &userv1.ListUserByIdReq{
			Id: userIds,
		})
		if err2 == nil {
			userMap = make(map[uint64]string, len(users.Items))
			for _, u := range users.Items {
				userMap[u.Id] = u.Name
			}
		}
		defer group.Done()
	}()
	go func() {
		foods, err2 := receiver.fc.ListByIds(ctx, &foodv1.ListFoodByIdReq{
			Id: foodIds,
		})
		if err2 == nil {
			foodMap = make(map[uint64]string, len(foods.Items))
			for _, f := range foods.Items {
				foodMap[f.Id] = f.Name
			}
		}
		defer group.Done()
	}()
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
	return result
}
