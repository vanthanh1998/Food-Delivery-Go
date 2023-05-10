package middleware

import (
	"Food-Delivery/common"
	"Food-Delivery/component/appctx"
	"Food-Delivery/component/tokenprovider/jwt"
	usermodel "Food-Delivery/module/user/model"
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"strings"
)

type AuthenticateStore interface {
	FindUser(ctx context.Context, conditions map[string]interface{}, moreInfo ...string) (*usermodel.User, error)
}

func ErrWrongAuthHeader(err error) *common.AppError {
	return common.NewCustomError(
		err,
		fmt.Sprintf("wrong authen header"),
		fmt.Sprintf("ErrWrongAuthHeader"),
	)
}

func extractTokenFromHeaderString(s string) (string, error) {
	parts := strings.Split(s, " ")
	// "Authorization" : "Bearer {token}"
	if parts[0] != "Bearer" || len(parts) < 2 || strings.TrimSpace(parts[1]) == "" {
		return "", ErrWrongAuthHeader(nil)
	}

	return parts[1], nil
}

// 1. Get token from header
// 2. Validate token and parse to payload
// 3. From the token payload, we use user_id find from db
func RequiredAuth(appCtx appctx.AppContext, authStore AuthenticateStore) gin.HandlerFunc {
	tokenProvider := jwt.NewTokenJWTProvider(appCtx.SecretKey())
	return func(c *gin.Context) {
		token, err := extractTokenFromHeaderString(c.GetHeader("Authorization"))

		if err != nil {
			panic(err)
		}

		//db := appCtx.GetMailDBConnection()
		//store := userstore.NewSQLStore(db)

		payload, err := tokenProvider.Validate(token)
		if err != nil {
			panic(err)
		}

		//user, err := store.FindUser(c.Request.Context(), map[string]interface{}{"id": payload.UserId})

		user, err := authStore.FindUser(c.Request.Context(), map[string]interface{}{"id": payload.UserId})

		if err != nil {
			panic(err)
		}

		if user.Status == 0 {
			panic(common.ErrNoPermission(errors.New("user has been deleted or banned")))
		}

		user.Mask(false)

		c.Set(common.CurrentUser, user) // c.Set("user", "value") để lấy ra dùng c.MustGet("user")
		c.Next()
	}

}
