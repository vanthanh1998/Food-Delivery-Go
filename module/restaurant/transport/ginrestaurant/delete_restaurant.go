package ginrestaurant

import (
	restaurantbiz "Food-Delivery/module/restaurant/biz"
	restaurantstorage "Food-Delivery/module/restaurant/storage"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

func DeleteRestaurant(db *gorm.DB) gin.HandlerFunc { // gin.HandlerFunc ~~ c *gin.Contex
	return func(c *gin.Context) {
		// get id use c.Params
		id, err := strconv.Atoi(c.Param("id")) // Atoi: chuyển đổi một chuỗi số sang dạng số nguyên (int)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		store := restaurantstorage.NewSQLStore(db) // get db
		biz := restaurantbiz.NewDeleteRestaurantBiz(store)

		if err := biz.DeleteRestaurant(c.Request.Context(), id); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":     err.Error(),
				"thanhrain": "cmm",
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"data": 1,
			})
		}
	}
}
