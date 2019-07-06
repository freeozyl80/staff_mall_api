package manage

import (
	"staff-mall-center/src/middleware/authcheck"
	"staff-mall-center/src/router/manage/admin"

	"github.com/gin-gonic/gin"
)

func MangeRouterInit(mgrouter *gin.RouterGroup) {
	mgrouter.Use(authcheck.AuthRequired)
	mgrouter.POST("/admin/login", admin.AdminLogin)
	mgrouter.POST("/admin/check", admin.AdminCheck)
}
