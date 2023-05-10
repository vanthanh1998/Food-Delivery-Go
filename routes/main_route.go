package routes

import (
	"Food-Delivery/component/appctx"
	"Food-Delivery/memcache"
	"Food-Delivery/middleware"
	"Food-Delivery/module/restaurant/transport/ginrestaurant"
	"Food-Delivery/module/restaurantlike/transport/ginrstlike"
	"Food-Delivery/module/upload/uploadtransport/ginupload"
	userstore "Food-Delivery/module/user/store"
	"Food-Delivery/module/user/transport/ginuser"
	"github.com/gin-gonic/gin"
)

func SetupRoute(appContext appctx.AppContext, v1 *gin.RouterGroup) {
	// call caching ~~ redis
	userStore := userstore.NewSQLStore(appContext.GetMailDBConnection())
	userCachingStore := memcache.NewUserCaching(memcache.NewCaching(), userStore)

	v1.POST("/upload", ginupload.Upload(appContext))

	// user
	v1.POST("/register", ginuser.Register(appContext))
	v1.POST("/authenticate/login", ginuser.Login(appContext))
	v1.GET("/profile", middleware.RequiredAuth(appContext, userCachingStore), ginuser.Profile(appContext))

	restaurants := v1.Group("/restaurants", middleware.RequiredAuth(appContext, userCachingStore))

	restaurants.GET("", ginrestaurant.ListRestaurant(appContext))
	restaurants.GET("/:id", ginrestaurant.GetIdRestaurant(appContext))
	restaurants.POST("", ginrestaurant.CreateRestaurant(appContext))
	restaurants.PUT("/:id", ginrestaurant.UpdateRestaurant(appContext))
	restaurants.DELETE("/:id", ginrestaurant.DeleteRestaurant(appContext))

	restaurants.POST("/:id/liked-users", ginrstlike.UserLikeRestaurant(appContext))
	restaurants.DELETE("/:id/liked-users", ginrstlike.UserDislikeRestaurant(appContext))
	restaurants.GET("/:id/liked-users", ginrstlike.ListUsers(appContext))
}
