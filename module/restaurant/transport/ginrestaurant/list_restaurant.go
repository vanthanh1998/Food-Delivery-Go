package ginrestaurant

import (
	"Food-Delivery/common"
	"Food-Delivery/component/appctx"
	restaurantbiz "Food-Delivery/module/restaurant/biz"
	restaurantmodel "Food-Delivery/module/restaurant/model"
	restaurantrepo "Food-Delivery/module/restaurant/repository"
	restaurantstorage "Food-Delivery/module/restaurant/storage"
	restaurantlikestore "Food-Delivery/module/restaurantlike/storage"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ListRestaurant(appCtx appctx.AppContext) gin.HandlerFunc { // gin.HandlerFunc ~~ c *gin.Contex
	return func(c *gin.Context) {
		db := appCtx.GetMailDBConnection() // database
		var data restaurantmodel.RestaurantCreate

		if err := c.ShouldBind(&data); err != nil {
			// ShouldBind: dùng để đọc và gán giá trị từ các request parameter vào các struct
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		var pagingData common.Paging

		if err := c.ShouldBind(&pagingData); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		pagingData.Fulfill()

		var filter restaurantmodel.Filter

		if err := c.ShouldBind(&filter); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		// db.Create(&data) ---- start ----
		store := restaurantstorage.NewSQLStore(db) // storage: call db
		likeStore := restaurantlikestore.NewSQLStore(db)
		repo := restaurantrepo.NewListRestaurantRepo(store, likeStore) // tầng repository => gánh phần liên kết dữ liệu
		biz := restaurantbiz.NewListRestaurantBiz(repo)

		result, err := biz.ListRestaurant(c.Request.Context(), &filter, &pagingData)

		if err != nil {
			panic(err) // panic tầng biz thì panic chính nó
		}
		// ---- end ----

		for i := range result {
			result[i].Mask(false)
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(result, pagingData, filter))

	}
}
