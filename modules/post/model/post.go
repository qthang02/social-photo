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
	Caption string `json:"caption" gorm:"column:caption;"`
}

func (Post) TableName() string { return "posts" }

type PostCreation struct {
	Id      int    `json:"-" gorm:"column:id;"`
	Caption string `json:"caption" binding:"required"`
}

func (PostCreation) TableName() string {
	return Post{}.TableName()
}

type PostUpdate struct {
	Caption *string `json:"caption"`
}

func (PostUpdate) TableName() string {
	return Post{}.TableName()
}
