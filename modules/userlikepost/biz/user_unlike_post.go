package biz

import (
	"context"
	"log"
	"social-photo/common"
	"social-photo/modules/userlikepost/model"
	"social-photo/pubsub"
)

type UserUnlikePostStore interface {
	Find(ctx context.Context, userId, postId int) (*model.Like, error)
	Delete(ctx context.Context, userId, postId int) error
}

type userUnlikePostBiz struct {
	store UserUnlikePostStore
	ps    pubsub.PubSub
}

func NewUserUnlikePostBiz(store UserUnlikePostStore, ps pubsub.PubSub) *userUnlikePostBiz {
	return &userUnlikePostBiz{store: store, ps: ps}
}

func (biz *userUnlikePostBiz) UnlikePost(ctx context.Context, userId, postId int) error {
	_, err := biz.store.Find(ctx, userId, postId)

	// Delete if data existed
	if err == common.RecordNotFound {
		return model.ErrDidNotLikeItem(err)
	}

	if err != nil {
		return model.ErrCannotUnlikeItem(err)
	}

	if err := biz.store.Delete(ctx, userId, postId); err != nil {
		return model.ErrCannotUnlikeItem(err)
	}

	if err := biz.ps.Publish(ctx, common.TopicUserUnlikedPost, pubsub.NewMessage(&model.Like{UserId: userId, PostId: postId})); err != nil {
		log.Println(err)
	}

	return nil
}
