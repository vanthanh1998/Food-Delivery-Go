package ginrestaurant

import (
	"Food-Delivery/common"
	"Food-Delivery/component/appctx"
	restaurantbiz "Food-Delivery/module/restaurant/biz"
	restaurantstorage "Food-Delivery/module/restaurant/storage"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func DeleteRestaurant(appCtx appctx.AppContext) gin.HandlerFunc { // gin.HandlerFunc ~~ c *gin.Contex
	return func(c *gin.Context) {
		db := appCtx.GetMailDBConnection() // database
		// get id use c.Params
		id, err := strconv.Atoi(c.Param("id")) // Atoi: chuyển đổi một chuỗi số sang dạng số nguyên (int)
		if err != nil {
			panic(common.ErrInvalidRequest(err)) // not panic(err)  bởi vì strconv.Atoi(c.Param("id")) là lỗi golang trả về
		}

		store := restaurantstorage.NewSQLStore(db) // get db
		biz := restaurantbiz.NewDeleteRestaurantBiz(store)

		if err := biz.DeleteRestaurant(c.Request.Context(), id); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
