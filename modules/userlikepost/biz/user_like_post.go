package biz

import (
	"context"
	"log"
	"social-photo/modules/userlikepost/model"
)

type UserLikePostStore interface {
	Create(ctx context.Context, data *model.Like) error
}

type IncreasePostStorage interface {
	IncreaseLikeCount(ctx context.Context, id int) error
}

type userLikePostBiz struct {
	store     UserLikePostStore
	postStore IncreasePostStorage
}

func NewUserLikePostBiz(store UserLikePostStore, postStorage IncreasePostStorage) *userLikePostBiz {
	return &userLikePostBiz{store: store, postStore: postStorage}
}

func (biz *userLikePostBiz) LikePost(ctx context.Context, data *model.Like) error {
	if err := biz.store.Create(ctx, data); err != nil {
		return model.ErrCannotLikeItem(err)
	}

	go func() {
		if err := biz.postStore.IncreaseLikeCount(ctx, data.PostId); err != nil {
			log.Fatalln("IncreaseLikeCount", err)
		}
	}()

	return nil
}
