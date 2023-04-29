package ginuser

import (
	"Food-Delivery/common"
	"Food-Delivery/component/appctx"
	"Food-Delivery/component/hasher"
	userbiz "Food-Delivery/module/user/biz"
	usermodel "Food-Delivery/module/user/model"
	userstore "Food-Delivery/module/user/store"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Register(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := appCtx.GetMailDBConnection()
		var data usermodel.UserCreate

		// khi check err mà call đến db thì bắt buộc phải có .Error
		if err := c.ShouldBind(&data); err != nil {
			panic(err)
		}
		// storage ~~ call db storage.go trong folder storage
		store := userstore.NewSQLStore(db)
		// md5: hash pw => call từ folder component/hasher
		md5 := hasher.NewMd5Hash()
		// biz new func
		biz := userbiz.NewRegisterBusiness(store, md5) // registerStore RegisterStore, hasher Hasher

		if err := biz.Register(c.Request.Context(), &data); err != nil {
			panic(err)
		}

		data.Mask(false)

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data.FakeId.String()))
	}
}
