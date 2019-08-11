package manage

import (
	"staff-mall-center/pkg/context"
	"staff-mall-center/src/middleware/authcheck"
	"staff-mall-center/src/router/manage/firm"
	"staff-mall-center/src/router/manage/user"

	"github.com/gin-gonic/gin"
)

func MangeRouterInit(mgrouter *gin.RouterGroup) {
	// 用户登录
	mgrouter.POST("/admin/login", context.Handle(user.AdminLogin))
	mgrouter.POST("/user/login", context.Handle(user.UserLogin))

	// access
	mgrouter.Use(context.Handle(authcheck.AuthRequired))
	mgrouter.GET("/admin/auth", context.Handle(user.AdminCheck))

	// 用户操作
	mgrouter.GET("/user/userinfo", context.Handle(user.AuthInfo))
	mgrouter.GET("/user/list", context.Handle(user.UserList))
	mgrouter.POST("/user/register", context.Handle(user.UserRegister))
	mgrouter.POST("/user/import", context.Handle(user.UserImport))

	// 公司列表操作
	mgrouter.GET("/firm/list", context.Handle(firm.FirmList))
	mgrouter.POST("/firm/add", context.Handle(firm.FirmAdd))

	// 商品操作

	//mgrouter.GET("/product/import", context.Handle(product.ProductImport))

}
