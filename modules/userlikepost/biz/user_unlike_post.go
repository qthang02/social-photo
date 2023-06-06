package biz

import (
	"context"
	"log"
	"social-photo/common"
	"social-photo/modules/userlikepost/model"
)

type UserUnlikePostStore interface {
	Find(ctx context.Context, userId, postId int) (*model.Like, error)
	Delete(ctx context.Context, userId, postId int) error
}

type DecreasePostStorage interface {
	DecreaseLikeCount(ctx context.Context, id int) error
}

type userUnlikePostBiz struct {
	store     UserUnlikePostStore
	postStore DecreasePostStorage
}

func NewUserUnlikePostBiz(store UserUnlikePostStore, postStorage DecreasePostStorage) *userUnlikePostBiz {
	return &userUnlikePostBiz{store: store, postStore: postStorage}
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

	go func() {
		if err := biz.postStore.DecreaseLikeCount(ctx, postId); err != nil {
			log.Fatalln("DecreaseLikeCount", err)
		}
	}()

	return nil
}
