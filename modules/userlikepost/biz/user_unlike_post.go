package biz

import (
	"context"
	"social-photo/common"
	"social-photo/modules/userlikepost/model"
)

type UserUnlikePostStore interface {
	Find(ctx context.Context, userId, postId int) (*model.Like, error)
	Delete(ctx context.Context, userId, postId int) error
}

type userUnlikePostBiz struct {
	store UserUnlikePostStore
}

func NewUserUnlikePostBiz(store UserUnlikePostStore) *userUnlikePostBiz {
	return &userUnlikePostBiz{store: store}
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

	return nil
}
