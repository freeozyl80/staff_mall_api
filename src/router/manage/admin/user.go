package admin

import "github.com/gin-gonic/gin"

func AdminLogin(ctx *gin.Context) {
	name := ctx.PostForm("name")
	pwd := ctx.PostForm("pwd")

	ctx.JSON(200, gin.H{
		"status": "posted",
		"name":   "登录成功" + name + ":" + pwd,
	})
}

func AdminCheck(ctx *gin.Context) {

}

func AdminRegister(ctx *gin.Context) {

}

func AddUser(ctx *gin.Context) {

}
