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

	mgrouter.POST("/user/register", context.Handle(user.UserRegister))

	// management 功能
	mgrouter.POST("/user/upload", context.Handle(user.UserImport))
}
