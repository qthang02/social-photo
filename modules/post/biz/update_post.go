package biz

import (
	"context"
	"errors"
	"social-photo/common"
	"social-photo/modules/post/model"
)

type UpdatePostStorage interface {
	GetPost(ctx context.Context, cond map[string]interface{}) (*model.Post, error)
	UpdatePost(ctx context.Context, cond map[string]interface{}, data *model.PostUpdate) error
}

type updatePostBiz struct {
	store     UpdatePostStorage
	requester common.Requester
}

func NewUpdatePostBiz(store UpdatePostStorage, requester common.Requester) *updatePostBiz {
	return &updatePostBiz{store: store, requester: requester}
}

func (biz *updatePostBiz) UpdatePostById(ctx context.Context, id int, dataUpdate *model.PostUpdate) error {
	data, err := biz.store.GetPost(ctx, map[string]interface{}{"id": id})

	if err != nil {
		if err == common.RecordNotFound {
			return common.ErrEntityNotFound(model.EntityName, err)
		}

		return common.ErrCannotUpdateEntity(model.EntityName, err)
	}

	isOwner := biz.requester.GetUserId() == data.UserId

	if !isOwner && !common.IsAdmin(biz.requester) {
		return common.ErrNoPermission(errors.New("you don't have permission to update this post"))
	}

	if err := biz.store.UpdatePost(ctx, map[string]interface{}{"id": id}, dataUpdate); err != nil {
		return common.ErrCannotUpdateEntity(model.EntityName, err)
	}

	return nil
}
