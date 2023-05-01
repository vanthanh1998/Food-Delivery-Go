package ginrstlike

import (
	"Food-Delivery/common"
	"Food-Delivery/component/appctx"
	rstlikebiz "Food-Delivery/module/restaurantlike/biz"
	restaurantlikemodel "Food-Delivery/module/restaurantlike/model"
	restaurantlikestorage "Food-Delivery/module/restaurantlike/storage"
	"github.com/gin-gonic/gin"
	"net/http"
)

// POST /v1/restaurants/:id/dislike
func UserDislikeRestaurant(appCtx appctx.AppContext) gin.HandlerFunc {
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

		store := restaurantlikestorage.NewSQLStore(db)
		//decStore := restaurantstorage.NewSQLStore(db)
		biz := rstlikebiz.NewUserDislikeRestaurantBiz(store, appCtx.GetPubSub())

		if err := biz.DislikeRestaurant(c.Request.Context(), &data); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
