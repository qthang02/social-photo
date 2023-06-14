package common

import "fmt"

const (
	CurrentUser = "current_user"
)

type DbType int

const (
	DbTypePost DbType = 1
	DbTypeUser DbType = 2

	TopicUserLikedPost   = "TopicUserLikedPost"
	TopicUserUnlikedPost = "TopicUserUnlikedPost"
)

func Recovery() {
	if r := recover(); r != nil {
		fmt.Println("Recovered:", r)
	}
}

type TokenPayload struct {
	UId   int    `json:"user_id"`
	URole string `json:"role"`
}

func (p TokenPayload) UserId() int {
	return p.UId
}

func (p TokenPayload) Role() string {
	return p.URole
}

type Requester interface {
	GetUserId() int
	GetRole() string
	GetEmail() string
}

func IsAdmin(requester Requester) bool {
	return requester.GetRole() == "admin" || requester.GetRole() == "mod"
}
