package third

import (
	"staff-mall-center/pkg/context"
	"staff-mall-center/src/middleware/authcheck"
	"staff-mall-center/src/router/third/qiniu"

	"github.com/gin-gonic/gin"
)

func ThirdRouterInit(trouter *gin.RouterGroup) {

	trouter.Use(context.Handle(authcheck.AuthRequired))
	trouter.POST("/qiniu/token", context.Handle(qiniu.GetToken))
}
