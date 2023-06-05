package biz

import (
	"context"
	"social-photo/common"
	"social-photo/modules/userlikepost/model"
)

type ListUserLikePostStore interface {
	ListUser(ctx context.Context, postId int, paging *common.Paging) ([]common.SimpleUser, error)
}

type listUserLikePostBiz struct {
	store ListUserLikePostStore
}

func NewListUserLikePostBiz(store ListUserLikePostStore) *listUserLikePostBiz {
	return &listUserLikePostBiz{store: store}
}

func (biz *listUserLikePostBiz) ListUserLikePost(ctx context.Context, postId int, paging *common.Paging) ([]common.SimpleUser, error) {
	result, err := biz.store.ListUser(ctx, postId, paging)
	if err != nil {
		return nil, common.ErrCannotListEntity(model.EntityName, err)
	}
	return result, nil
}
