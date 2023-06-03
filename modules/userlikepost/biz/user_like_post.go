package biz

import (
	"context"
	"social-photo/modules/userlikepost/model"
)

type UserLikePostStore interface {
	Create(ctx context.Context, data *model.Like) error
}

type userLikePostBiz struct {
	store UserLikePostStore
}

func NewUserLikePostBiz(store UserLikePostStore) *userLikePostBiz {
	return &userLikePostBiz{store: store}
}

func (biz *userLikePostBiz) LikePost(ctx context.Context, data *model.Like) error {
	if err := biz.store.Create(ctx, data); err != nil {
		return model.ErrCannotLikeItem(err)
	}
	return nil
}
