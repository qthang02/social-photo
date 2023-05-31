package model

import (
	"errors"
	"social-photo/common"
)

const (
	EntityName = "Post"
)

var (
	ErrCaptionIsBlank = errors.New("Caption is blank")
)

type Post struct {
	common.SQLModel
	UserId  int                `json:"user_id" gorm:"column:user_id;"`
	Caption string             `json:"caption" gorm:"column:caption;"`
	Image   *common.Images     `json:"image" gorm:"column:image;"`
	Owner   *common.SimpleUser `json:"owner" gorm:"foreignKey:UserId;"`
}

func (Post) TableName() string { return "posts" }

type PostCreation struct {
	Id      int            `json:"-" gorm:"column:id;"`
	UserId  int            `json:"-" gorm:"column:user_id;"`
	Caption string         `json:"caption" binding:"required"`
	Image   *common.Images `json:"image" gorm:"column:image;"`
}

func (PostCreation) TableName() string {
	return Post{}.TableName()
}

type PostUpdate struct {
	Caption *string `json:"caption"`
	Image   *string `json:"image"`
}

func (PostUpdate) TableName() string {
	return Post{}.TableName()
}
