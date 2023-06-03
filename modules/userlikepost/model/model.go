package model

import (
	"fmt"
	"social-photo/common"
	"time"
)

const (
	EntityName = "UserLikePost"
)

type Like struct {
	UserId    int                `json:"user_id" gorm:"column:user_id;"`
	PostId    int                `json:"post_id" gorm:"column:post_id;"`
	CreatedAt *time.Time         `json:"created_at" gorm:"column:created_at;"`
	User      *common.SimpleUser `json:"-" gorm:"foreignKey:UserId;"`
}

func (Like) TableName() string { return "user_like_posts" }

func ErrCannotLikeItem(err error) *common.AppError {
	return common.NewCustomError(
		err,
		fmt.Sprintf("Cannot like this item"),
		fmt.Sprintf("ErrCannotLikeItem"),
	)
}

func ErrCannotUnlikeItem(err error) *common.AppError {
	return common.NewCustomError(
		err,
		fmt.Sprintf("Cannot dislike this item"),
		fmt.Sprintf("ErrCannotUnlikeItem"),
	)
}

func ErrDidNotLikeItem(err error) *common.AppError {
	return common.NewCustomError(
		err,
		fmt.Sprintf("You have not liked this item"),
		fmt.Sprintf("ErrCannotDidNotLikeItem"),
	)
}
