package manage

import (
	"staff-mall-center/pkg/context"
	"staff-mall-center/src/middleware/authcheck"
	"staff-mall-center/src/router/manage/user"

	"github.com/gin-gonic/gin"
)

func MangeRouterInit(mgrouter *gin.RouterGroup) {
	// 用户组
	mgrouter.POST("/admin/login", context.Handle(user.AdminLogin))
	mgrouter.POST("/user/login", context.Handle(user.UserLogin))

	mgrouter.Use(context.Handle(authcheck.AuthRequired))

	// access
	mgrouter.GET("/admin/auth", context.Handle(user.AdminCheck))

	mgrouter.GET("/user/list", context.Handle(user.UserList))
	mgrouter.POST("/user/register", context.Handle(user.UserRegister))
	mgrouter.POST("/user/import", context.Handle(user.UserImport))
}
