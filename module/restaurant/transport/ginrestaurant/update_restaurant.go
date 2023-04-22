package ginrestaurant

import (
	"Food-Delivery/component/appctx"
	restaurantbiz "Food-Delivery/module/restaurant/biz"
	restaurantmodel "Food-Delivery/module/restaurant/model"
	restaurantstorage "Food-Delivery/module/restaurant/storage"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func UpdateRestaurant(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id")) // Atoi: string -> (int)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		var data restaurantmodel.RestaurantUpdate
		db := appCtx.GetMailDBConnection()

		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		// store
		store := restaurantstorage.NewSQLStore(db)
		// biz
		biz := restaurantbiz.NewUpdateRestaurantBiz(store) // nếu k có hàm Updateở trong storeafe/update.go thì sẽ báo err
		// xử lý error
		if err := biz.UpdateRestaurant(c.Request.Context(), id, data); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, gin.H{
			"status": 1,
			"data":   data,
		})
	}
}
