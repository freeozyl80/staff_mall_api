package router

import (
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"staff-mall-center/pkg/setting"
	"staff-mall-center/src/router/benchmark"
	"staff-mall-center/src/router/manage"
	"staff-mall-center/src/router/wx"
	"syscall"

	"github.com/fvbock/endless"
	//"github.com/gin-contrib/cors"

	"github.com/gin-contrib/static"
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

	// cors
	// router.Use(cors.New(cors.Config{
	// 	AllowOrigins:     []string{"http://127.0.0.1:8001"},
	// 	AllowMethods:     []string{"GET", "POST", "OPTIONS"},
	// 	AllowHeaders:     []string{"Access-Control-Allow-Headers, Origin,Accept, X-Requested-With, Content-Type, Access-Control-Request-Method, Access-Control-Request-Headers, hualvmall_authorization"},
	// 	ExposeHeaders:    []string{"Content-Length"},
	// 	AllowCredentials: true,
	// 	AllowOriginFunc: func(origin string) bool {
	// 		return origin == "http://hualvmall.com"
	// 	},
	// 	MaxAge: 12 * time.Hour,
	// }))

	// 心跳benchmark检测
	router.GET("/benchmark", benchmark.MyBenchLogger, func(c *gin.Context) {
		resp := map[string]string{"hello": "world"}

		c.JSON(200, resp)

	})
	// 	暂定router 分为三种类型 分别是 common || wx相关 || mis相关，具体后面再细分吧
	path, _ := filepath.Abs("./www/index.html")
	router.LoadHTMLFiles(path)

	router.Use(static.Serve("/web", static.LocalFile("./www", false)))

	router.GET("/web/*action", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})

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
