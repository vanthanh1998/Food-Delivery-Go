package ginuser

import (
	"Food-Delivery/common"
	"Food-Delivery/component/appctx"
	"Food-Delivery/component/hasher"
	"Food-Delivery/component/tokenprovider/jwt"
	userbiz "Food-Delivery/module/user/biz"
	usermodel "Food-Delivery/module/user/model"
	userstore "Food-Delivery/module/user/store"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Login(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := appCtx.GetMailDBConnection()
		var dataUserData usermodel.UserLogin

		if err := c.ShouldBind(&dataUserData); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		tokenProvider := jwt.NewTokenJWTProvider(appCtx.SecretKey())
		// storage
		store := userstore.NewSQLStore(db)
		// md5
		md5 := hasher.NewMd5Hash()
		// biz
		business := userbiz.NewLoginBusiness(store, tokenProvider, md5, 60*60*24*30)

		account, err := business.Login(c.Request.Context(), &dataUserData)

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(account))
	}
}
