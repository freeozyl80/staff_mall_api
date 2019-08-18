package wx

import (
	"staff-mall-center/pkg/context"
	"staff-mall-center/src/router/wx/product"

	"github.com/gin-gonic/gin"
)

func WxRouterInit(wrouter *gin.RouterGroup) {

	wrouter.GET("/product/firm/list", context.Handle(product.ProductFirmList))
	wrouter.GET("/category/firm/list", context.Handle(product.CategoryFirmList))
}
