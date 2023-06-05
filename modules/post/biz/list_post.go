package biz

import (
	"golang.org/x/net/context"
	"social-photo/common"
	"social-photo/modules/post/model"
)

type ListPostRepo interface {
	ListPost(ctx context.Context, paging *common.Paging, moreKey ...string) ([]model.Post, error)
}

type listPostBiz struct {
	Repo      ListPostRepo
	requester common.Requester
}

func NewListPostBiz(repo ListPostRepo, requester common.Requester) *listPostBiz {
	return &listPostBiz{Repo: repo, requester: requester}
}

func (biz *listPostBiz) ListPostBiz(ctx context.Context, paging *common.Paging) ([]model.Post, error) {
	data, err := biz.Repo.ListPost(ctx, paging, "Owner")

	if err != nil {
		return nil, common.ErrCannotListEntity(model.EntityName, err)
	}

	return data, nil
}
