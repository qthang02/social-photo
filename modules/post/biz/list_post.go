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
	store     ListPostStorage
	requester common.Requester
}

func NewListPostBiz(store ListPostStorage, requester common.Requester) *listPostBiz {
	return &listPostBiz{store: store, requester: requester}
}

func (biz *listPostBiz) ListPostBiz(ctx context.Context, paging *common.Paging) ([]model.Post, error) {
	ctxStore := context.WithValue(ctx, common.CurrentUser, biz.requester)

	data, err := biz.store.ListPost(ctxStore, paging)

	if err != nil {
		return nil, common.ErrCannotListEntity(model.EntityName, err)
	}

	return data, nil
}
