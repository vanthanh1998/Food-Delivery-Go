package ginrstlike

import (
	"Food-Delivery/common"
	"Food-Delivery/component/appctx"
	rstlikebiz "Food-Delivery/module/restaurantlike/biz"
	restaurantlikemodel "Food-Delivery/module/restaurantlike/model"
	restaurantlikestore "Food-Delivery/module/restaurantlike/storage"
	"github.com/gin-gonic/gin"
	"net/http"
)

// POST /v1/restaurants/:id/like
func UserLikeRestaurant(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := appCtx.GetMailDBConnection()

		uid, err := common.FromBase58(c.Param("id"))

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		requester := c.MustGet(common.CurrentUser).(common.Requester)

		data := restaurantlikemodel.Like{
			RestaurantId: int(uid.GetLocalID()),
			UserId:       requester.GetUserId(),
		}

		store := restaurantlikestore.NewSQLStore(db)
		biz := rstlikebiz.NewUserLikeRestaurantBiz(store)

		if err := biz.UserLikeRestaurant(c.Request.Context(), &data); err != nil {
			c.JSON(http.StatusOK, common.SimpleSuccessResponse("like ok"))
		}
	}
}