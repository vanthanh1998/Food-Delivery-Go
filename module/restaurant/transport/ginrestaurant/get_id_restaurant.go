package ginrestaurant

import (
	"Food-Delivery/common"
	"Food-Delivery/component/appctx"
	restaurantbiz "Food-Delivery/module/restaurant/biz"
	restaurantmodel "Food-Delivery/module/restaurant/model"
	restaurantstorage "Food-Delivery/module/restaurant/storage"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetIdRestaurant(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		var data restaurantmodel.Restaurant
		db := appCtx.GetMailDBConnection()
		store := restaurantstorage.NewSQLStore(db)
		biz := restaurantbiz.NewGetIdRestaurantBiz(store)

		if err := biz.GetIdRestaurant(c.Request.Context(), id, &data); err != nil {
			panic(err)
		}
		data.Mask(false)

		c.JSON(http.StatusOK, gin.H{
			"data": data,
		})
	}
}
