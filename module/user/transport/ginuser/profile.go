package ginuser

import (
	"Food-Delivery/common"
	"Food-Delivery/component/appctx"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Profile(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		u := c.MustGet(common.CurrentUser).(common.Requester) // ép kiểu (type assertion) => (common.Requester)
		//data := u.GetEmail()
		//jsonData, _ := json.MarshalIndent(data, "", "  ")
		//log.Println((string(jsonData)))

		// Làm cho th update password
		//newPW := "abcxyz"
		//type update struct {
		//	NewPass *string
		//}
		//log.Println(update{NewPass: &newPW})

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(u))
	}
}
