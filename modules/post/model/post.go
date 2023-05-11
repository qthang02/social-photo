package model

import (
	"errors"
	"time"
)

var (
	ErrCaptionIsBlank = errors.New("Caption is blank")
)

type Post struct {
	Id        int        `json:"id" gorm:"column:id;"`
	Caption   string     `json:"caption" gorm:"column:caption;"`
	CreatedAt *time.Time `json:"created_at" gorm:"column:created_at;"`
	UpdatedAt *time.Time `json:"updated_at" gorm:"column:updated_at;"`
}

func (Post) TableName() string { return "posts" }

type PostCreation struct {
	Id      int    `json:"-" gorm:"column:id;"`
	Caption string `json:"caption" binding:"required"`
}

func (PostCreation) TableName() string {
	return Post{}.TableName()
}
