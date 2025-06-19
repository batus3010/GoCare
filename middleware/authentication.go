package middleware

import (
	"GoCare/common"
	"GoCare/components/appctx"
	"GoCare/components/tokenprovider/jwt"
	userModel "GoCare/module/user/model"
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"strings"
)

type AuthenStore interface {
	FindUser(ctx context.Context, conditions map[string]interface{}) (*userModel.User, error)
}

func ErrWrongAuthHeader(err error) *common.AppError {
	return common.NewCustomErrorResponse(
		err,
		fmt.Sprintf("wrong authen header"),
		fmt.Sprintf("ErrWrongAuthHeader"),
	)
}

func extractTokenFromHeaderString(s string) (string, error) {
	parts := strings.Split(s, " ")

	if parts[0] != "Bearer" || len(parts) != 2 || strings.TrimSpace(parts[1]) == "" {
		return "", ErrWrongAuthHeader(nil)
	}
	return parts[1], nil
}

// 1. Get token from header
// 2. Validate token and parse to payload
// 3. From the token payload, we use user id to find from DB
func RequiredAuth(appCtx appctx.AppContext, authStore AuthenStore) func(c *gin.Context) {
	tokenProvider := jwt.NewTokenJwtProvider(appCtx.SecretKey())

	return func(c *gin.Context) {
		token, err := extractTokenFromHeaderString(c.Request.Header.Get("Authorization"))

		if err != nil {
			panic(err)
		}
		payload, err := tokenProvider.Validate(token)
		if err != nil {
			panic(err)
		}

		user, err := authStore.FindUser(c.Request.Context(), map[string]interface{}{"id": payload.UserId})

		if err != nil {
			panic(err)
		}

		if user.Status == 0 {
			panic(common.ErrorNoPermission(errors.New("user has been deleted")))
		}

		c.Set(common.CurrentUser, user)
		c.Next()
	}
}
