package router

import (
	"fmt"
	"log"
	"staff-mall-center/pkg/setting"
	"staff-mall-center/src/router/benchmark"
	"staff-mall-center/src/router/manage"
	"staff-mall-center/src/router/wx"
	"syscall"

	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
)

func Start() {
	router := gin.Default()
	// 	mode模式
	gin.SetMode(gin.DebugMode)

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
	// 	暂定router 分为三种类型 分别是 common || wx相关 || mis相关，具体后面再细分吧

	wxRouter := router.Group("/wx")
	wx.WxRouterInit(wxRouter)

	manageRouter := router.Group("/manage")
	manage.MangeRouterInit(manageRouter)

	port := fmt.Sprintf(":%d", setting.ServerSetting.HttpPort)
	//router.Run(port)

	// Test

	server := endless.NewServer(port, router)

	server.BeforeBegin = func(add string) {
		log.Printf("pid is %d", syscall.Getpid())
	}

	endless.DefaultReadTimeOut = setting.ServerSetting.ReadTimeout
	endless.DefaultWriteTimeOut = setting.ServerSetting.WriteTimeout
	endless.DefaultMaxHeaderBytes = 1 << 20

	err := server.ListenAndServe()
	if err != nil {
		log.Printf("Server err: %v", err)
	}
}