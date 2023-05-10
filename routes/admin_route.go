package routes

import (
	"Food-Delivery/component/appctx"
	"github.com/gin-gonic/gin"
)

func SetupAdminRoute(appContext appctx.AppContext, v1 *gin.RouterGroup) {
	//admin := v1.Group("/admin",
	//	middleware.RequiredAuth(appContext),
	//	middleware.RoleRequired(appContext, "admin", "mod"),
	//)
	//
	//{
	//	admin.GET("/profile", ginuser.Profile(appContext))
	//}
}
