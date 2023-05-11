package biz

import (
	"context"
	"social-photo/modules/post/model"
)

type DeletePostStorage interface {
	GetPost(ctx context.Context, cond map[string]interface{}) (*model.Post, error)
	DeletePost(ctx context.Context, cond map[string]interface{}) error
}

type deletePostBiz struct {
	store DeletePostStorage
}

func NewDeletePostBiz(store DeletePostStorage) *deletePostBiz {
	return &deletePostBiz{store: store}
}

func (biz *deletePostBiz) DeletePostById(ctx context.Context, id int) error {
	_, err := biz.store.GetPost(ctx, map[string]interface{}{"id": id})

	if err != nil {
		return err
	}

	if err := biz.store.DeletePost(ctx, map[string]interface{}{"id": id}); err != nil {
		return err
	}

	return nil
}
