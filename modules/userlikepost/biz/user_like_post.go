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

type userLikePostBiz struct {
	store UserLikePostStore
	ps    pubsub.PubSub
}

func NewUserLikePostBiz(store UserLikePostStore, ps pubsub.PubSub) *userLikePostBiz {
	return &userLikePostBiz{store: store, ps: ps}
}

func (biz *userLikePostBiz) LikePost(ctx context.Context, data *model.Like) error {
	if err := biz.store.Create(ctx, data); err != nil {
		return model.ErrCannotLikeItem(err)
	}

	if err := biz.ps.Publish(ctx, common.TopicUserLikedPost, pubsub.NewMessage(data)); err != nil {
		log.Println(err)
	}

	return nil
}
