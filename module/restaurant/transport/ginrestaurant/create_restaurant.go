package ginrestaurant

import (
	"Food-Delivery/common"
	"Food-Delivery/component/appctx"
	restaurantbiz "Food-Delivery/module/restaurant/biz"
	restaurantmodel "Food-Delivery/module/restaurant/model"
	restaurantstorage "Food-Delivery/module/restaurant/storage"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateRestaurant(appCtx appctx.AppContext) gin.HandlerFunc { // gin.HandlerFunc ~~ c *gin.Contex
	return func(c *gin.Context) {
		db := appCtx.GetMailDBConnection() // database

		requester := c.MustGet(common.CurrentUser).(common.Requester)

		//go func() {
		//	defer common.AppRecover()
		//
		//	arr := []int{}
		//	log.Println(arr[0])
		//}()

		var data restaurantmodel.RestaurantCreate

		if err := c.ShouldBind(&data); err != nil {
			// ShouldBind: dùng để đọc và gán giá trị từ các request parameter vào các struct
			panic(err) // Lưu ý: hàm này chỉ dùng ở tầng ngoài cùng (uploadtransport)
		}

		data.UserId = requester.GetUserId()

		// db.Create(&data) ---- start ----
		store := restaurantstorage.NewSQLStore(db) // store: call db
		biz := restaurantbiz.NewCreateRestaurantBiz(store)

		if err := biz.CreateRestaurant(c.Request.Context(), &data); err != nil {
			panic(err)
		}
		// ---- end ----

		data.Mask(false)

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data.FakeId.String()))
	}
}
