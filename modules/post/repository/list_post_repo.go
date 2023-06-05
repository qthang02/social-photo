package repository

import (
	"context"
	"social-photo/common"
	"social-photo/modules/post/model"
)

type ListPostStorage interface {
	ListPost(ctx context.Context, paging *common.Paging, moreKey ...string) ([]model.Post, error)
}

type PostLikeStorage interface {
	GetItemLikes(ctx context.Context, ids []int) (map[int]int, error)
}

type listPostRepo struct {
	store     ListPostStorage
	likeStore PostLikeStorage
	requester common.Requester
}

func NewListPostRepo(store ListPostStorage, likeStore PostLikeStorage, requester common.Requester) *listPostRepo {
	return &listPostRepo{store: store, requester: requester, likeStore: likeStore}
}

func (repo *listPostRepo) ListPost(ctx context.Context, paging *common.Paging, moreKey ...string) ([]model.Post, error) {
	data, err := repo.store.ListPost(ctx, paging, moreKey...)

	if err != nil {
		return nil, common.ErrCannotListEntity(model.EntityName, err)
	}

	if len(data) == 0 {
		return data, nil
	}

	ids := make([]int, len(data))

	for i := range ids {
		ids[i] = data[i].Id
	}

	likeUserMap, err := repo.likeStore.GetItemLikes(ctx, ids)
	if err != nil {
		return data, nil
	}

	for i := range data {
		data[i].LikeCount = likeUserMap[data[i].Id]
	}

	return data, nil
}
