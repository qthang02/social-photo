package biz

import (
	"context"
	"social-photo/common"
	"social-photo/modules/post/model"
	"strings"
)

type CreatePostStorage interface {
	CreatePost(ctx context.Context, data *model.PostCreation) error
}

type createPostBiz struct {
	store CreatePostStorage
}

func NewCreatePostBiz(store CreatePostStorage) *createPostBiz {
	return &createPostBiz{store: store}
}

func (biz *createPostBiz) CreateNewPost(ctx context.Context, data *model.PostCreation) error {

	caption := strings.TrimSpace(data.Caption)

	if caption == "" {
		return model.ErrCaptionIsBlank
	}

	if err := biz.store.CreatePost(ctx, data); err != nil {
		return common.ErrCannotCreateEntity(model.EntityName, err)
	}

	return nil
}
