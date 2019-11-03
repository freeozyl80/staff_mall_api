package manage

import (
	"staff-mall-center/pkg/context"
	"staff-mall-center/src/middleware/authcheck"
	"staff-mall-center/src/router/manage/firm"
	"staff-mall-center/src/router/manage/order"
	"staff-mall-center/src/router/manage/product"
	"staff-mall-center/src/router/manage/staff"
	"staff-mall-center/src/router/manage/user"

	"github.com/gin-gonic/gin"
)

func MangeRouterInit(mgrouter *gin.RouterGroup) {
	// 用户登录
	mgrouter.POST("/admin/login", context.Handle(user.AdminLogin))

	// access
	mgrouter.Use(context.Handle(authcheck.AuthRequired))
	mgrouter.GET("/admin/auth", context.Handle(user.AdminCheck))

	// 用户操作
	mgrouter.GET("/user/userinfo", context.Handle(user.AuthInfo))
	mgrouter.GET("/user/list", context.Handle(user.UserList))
	mgrouter.POST("/user/register", context.Handle(user.UserRegister))
	mgrouter.POST("/user/import", context.Handle(user.UserImport))

	mgrouter.GET("/firm/staff", context.Handle(user.StaffInfo))
	mgrouter.POST("/firm/staff/coin", context.Handle(staff.UpdateStaffCoin))
	mgrouter.POST("/firm/staff/list/coin", context.Handle(staff.UpdateStaffListCoin))

	// 公司列表操作
	mgrouter.GET("/firm/detail", context.Handle(firm.FirmDetail))
	mgrouter.GET("/firm/list", context.Handle(firm.FirmList))
	mgrouter.POST("/firm/add", context.Handle(firm.FirmAdd))
	mgrouter.POST("/firm/update", context.Handle(firm.FirmUpdate))
	//// 公司列表账户操作
	mgrouter.GET("/firm/account", context.Handle(firm.AccountList))
	mgrouter.GET("/firm/delegate", context.Handle(user.DelegateManager))
	//// 公司列表商品操作

	mgrouter.POST("/product/import", context.Handle(product.ProductImport))
	mgrouter.GET("/product/list", context.Handle(product.ProductList))
	mgrouter.GET("/product/firm/categroy/list", context.Handle(product.CategroyList))
	mgrouter.GET("/supplier/list", context.Handle(product.SupplierList))
	mgrouter.GET("/product/firm/list", context.Handle(product.ProductFirmList))
	mgrouter.POST("/product/firm/new", context.Handle(product.ProductFirmAdd))
	mgrouter.POST("/product/firm/update", context.Handle(product.ProductFirmUpdate))
	mgrouter.GET("/product/firm/detail", context.Handle(product.ProductFirmDetail))
	//// 公司列表订单操作
	mgrouter.GET("/order/list", context.Handle(order.OrderList))
	mgrouter.GET("/order/cancel", context.Handle(order.CancelOrder))
}
