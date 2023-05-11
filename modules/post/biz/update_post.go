package biz

import (
	"context"
	"social-photo/modules/post/model"
)

type UpdatePostStorage interface {
	GetPost(ctx context.Context, cond map[string]interface{}) (*model.Post, error)
	UpdatePost(ctx context.Context, cond map[string]interface{}, data *model.PostUpdate) error
}

type updatePostBiz struct {
	store UpdatePostStorage
}

func NewUpdatePostBiz(store UpdatePostStorage) *updatePostBiz {
	return &updatePostBiz{store: store}
}

func (biz updatePostBiz) UpdatePostById(ctx context.Context, id int, dataUpdate *model.PostUpdate) error {
	_, err := biz.store.GetPost(ctx, map[string]interface{}{"id": id})

	if err != nil {
		return err
	}

	if err := biz.store.UpdatePost(ctx, map[string]interface{}{"id": id}, dataUpdate); err != nil {
		return err
	}

	return nil
}
