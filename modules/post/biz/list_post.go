package biz

import (
	"golang.org/x/net/context"
	"social-photo/common"
	"social-photo/modules/post/model"
)

type ListPostStorage interface {
	ListPost(ctx context.Context, paging *common.Paging) ([]model.Post, error)
}

type listPostBiz struct {
	store ListPostStorage
}

func NewListPostBiz(store ListPostStorage) *listPostBiz {
	return &listPostBiz{store: store}
}

func (biz *listPostBiz) ListPostBiz(ctx context.Context, paging *common.Paging) ([]model.Post, error) {
	data, err := biz.store.ListPost(ctx, paging)

	if err != nil {
		return nil, err
	}

	return data, nil
}
