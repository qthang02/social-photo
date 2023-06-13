package biz

import (
	"context"
	"log"
	"social-photo/common"
	"social-photo/modules/userlikepost/model"
	"social-photo/pubsub"
)

type UserLikePostStore interface {
	Create(ctx context.Context, data *model.Like) error
}

//type IncreasePostStorage interface {
//	IncreaseLikeCount(ctx context.Context, id int) error
//}

type userLikePostBiz struct {
	store UserLikePostStore
	//postStore IncreasePostStorage
	ps pubsub.PubSub
}

func NewUserLikePostBiz(store UserLikePostStore,
	//postStorage IncreasePostStorage,
	ps pubsub.PubSub) *userLikePostBiz {
	return &userLikePostBiz{store: store,
		//postStore: postStorage,
		ps: ps}
}

func (biz *userLikePostBiz) LikePost(ctx context.Context, data *model.Like) error {
	if err := biz.store.Create(ctx, data); err != nil {
		return model.ErrCannotLikeItem(err)
	}

	//go func() {
	//	if err := biz.postStore.IncreaseLikeCount(ctx, data.PostId); err != nil {
	//		log.Fatalln("IncreaseLikeCount", err)
	//	}
	//}()

	if err := biz.ps.Publish(ctx, common.TopicUserLikedPost, pubsub.NewMessage(data)); err != nil {
		log.Println(err)
	}

	return nil
}
