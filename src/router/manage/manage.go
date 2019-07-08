package manage

import (
	"staff-mall-center/src/middleware/authcheck"
	"staff-mall-center/src/router/manage/user"

	"github.com/gin-gonic/gin"
)

func MangeRouterInit(mgrouter *gin.RouterGroup) {
	mgrouter.Use(authcheck.AuthRequired)
	mgrouter.POST("/user/register", user.UserRegister)
	mgrouter.POST("/admin/login", user.AdminLogin)
	mgrouter.POST("/admin/check", user.AdminCheck)
}
