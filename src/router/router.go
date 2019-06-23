package router

import (
	"fmt"
	"staff-mall-center/pkg/setting"
	"staff-mall-center/src/middleware/admin"
	"staff-mall-center/src/middleware/benchmark"
	r "staff-mall-center/src/router/admin"

	"github.com/gin-gonic/gin"
)

func Start() {
	router := gin.Default()
	// 	mode模式
	if setting.DevMode {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
	// prd模式开启log日志
	router.Use(gin.Logger())

	// 重启模式，panic下 500
	router.Use(gin.Recovery())

	// TODO中间件 的 genres 需要一个包裹，model还是pkg再想想
	// 心跳benchmark检测
	router.GET("/benchmark", benchmark.MyBenchLogger, func(c *gin.Context) {
		resp := map[string]string{"hello": "world"}

		c.JSON(200, resp)

	})

	authorized := router.Group("/")

	authorized.Use(admin.AuthRequired)
	{
		authorized.POST("/login", r.AdminLogin)
	}

	port := fmt.Sprintf(":%d", setting.Port)
	router.Run(port)
}
