package wx

import (
	"staff-mall-center/pkg/context"
	"staff-mall-center/src/middleware/authcheck"
	"staff-mall-center/src/router/manage/user"
	"staff-mall-center/src/router/wx/order"
	"staff-mall-center/src/router/wx/product"
	"staff-mall-center/src/router/wx/staff"

	"github.com/gin-gonic/gin"
)

func WxRouterInit(wrouter *gin.RouterGroup) {

	wrouter.GET("/product/firm/list", context.Handle(product.ProductFirmList))
	wrouter.GET("/category/firm/list", context.Handle(product.CategoryFirmList))

	wrouter.GET("/user/:uid", context.Handle(staff.UserInfo))

	wrouter.POST("/user/login", context.Handle(user.UserLogin))

	wrouter.Use(context.Handle(authcheck.StaffRequire))
	wrouter.POST("/product/order", context.Handle(order.GenerateOrder))
	wrouter.GET("/order/list", context.Handle(order.ListOrder))
}
