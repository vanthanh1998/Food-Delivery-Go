package routes

import (
	"Food-Delivery/component/appctx"
	"Food-Delivery/middleware"
	"Food-Delivery/module/restaurant/transport/ginrestaurant"
	"Food-Delivery/module/restaurantlike/transport/ginrstlike"
	"Food-Delivery/module/upload/uploadtransport/ginupload"
	"Food-Delivery/module/user/transport/ginuser"
	"github.com/gin-gonic/gin"
)

func SetupRoute(appContext appctx.AppContext, v1 *gin.RouterGroup) {
	v1.POST("/upload", ginupload.Upload(appContext))

	// user
	v1.POST("/register", ginuser.Register(appContext))
	v1.POST("/authenticate/login", ginuser.Login(appContext))
	v1.GET("/profile", middleware.RequiredAuth(appContext), ginuser.Profile(appContext))

	restaurants := v1.Group("/restaurants", middleware.RequiredAuth(appContext))

	restaurants.GET("", ginrestaurant.ListRestaurant(appContext))
	restaurants.GET("/:id", ginrestaurant.GetIdRestaurant(appContext))
	restaurants.POST("", ginrestaurant.CreateRestaurant(appContext))
	restaurants.PUT("/:id", ginrestaurant.UpdateRestaurant(appContext))
	restaurants.DELETE("/:id", ginrestaurant.DeleteRestaurant(appContext))

	restaurants.POST("/:id/like", ginrstlike.UserLikeRestaurant(appContext))
}
