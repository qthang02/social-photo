package middleware

import (
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"social-photo/common"
	"social-photo/component/tokenprovider"
	"social-photo/modules/user/model"
	"strings"
)

type AuthStore interface {
	FindUser(ctx context.Context, conditions map[string]interface{}) (*model.User, error)
}

func ErrWrongAuthHeader(err error) *common.AppError {
	return common.NewCustomError(
		err,
		fmt.Sprintf("Wrong authentication header"),
		fmt.Sprintf("ErrWrongAuthHeader"),
	)
}

func extractTokenFromHeader(s string) (string, error) {
	parts := strings.Split(s, " ")

	// "Authorization": "Bearer {token}"

	if parts[0] != "Bearer" || len(parts) < 2 || strings.TrimSpace(parts[1]) == "" {
		return "", ErrWrongAuthHeader(nil)
	}

	return parts[1], nil
}

func RequiredAuth(authStore AuthStore, tokenProvider tokenprovider.Provider) func(*gin.Context) {
	return func(c *gin.Context) {
		// 1. get token from header
		token, err := extractTokenFromHeader(c.GetHeader("Authorization"))

		if err != nil {
			panic(err)
		}

		// 2. validate token and parse to payload
		payload, err := tokenProvider.Validate(token)

		if err != nil {
			panic(err)
		}

		// 3. from the token payload, we use user_id to find user in db
		user, err := authStore.FindUser(c.Request.Context(), map[string]interface{}{"id": payload.UserId()})

		if err != nil {
			panic(err)
		}

		// 3.1. check if user is active
		if user.Status == 0 {
			panic(common.ErrNoPermission(errors.New("user is not active")))
		}

		// 4. set user to context
		c.Set(common.CurrentUser, user)

		c.Next()
	}
}
