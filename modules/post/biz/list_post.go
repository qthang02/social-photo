package biz

import (
	"golang.org/x/net/context"
	"social-photo/common"
	"social-photo/modules/post/model"
)

type ListPostStorage interface {
	ListPost(ctx context.Context, paging *common.Paging, moreKey ...string) ([]model.Post, error)
}

type listPostBiz struct {
	store     ListPostStorage
	requester common.Requester
}

func NewListPostBiz(store ListPostStorage, requester common.Requester) *listPostBiz {
	return &listPostBiz{store: store, requester: requester}
}

func (biz *listPostBiz) ListPostBiz(ctx context.Context, paging *common.Paging) ([]model.Post, error) {
	data, err := biz.store.ListPost(ctx, paging, "Owner")

	if err != nil {
		return nil, common.ErrCannotListEntity(model.EntityName, err)
	}

	return data, nil
}
