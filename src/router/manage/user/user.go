package user

import (
	"staff-mall-center/models/service/manage_service"

	"github.com/gin-gonic/gin"
)

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

func UserRegister(ctx *gin.Context) {
	username := ctx.Query("username")
	password := ctx.Query("password")

	user_service := manage_service.User{Usertype: 4, Username: username, Password: password}
	isExist, err := user_service.Register()

	if err != nil {
		ctx.JSON(200, gin.H{
			"status": "posted",
			"name":   "注册失败",
		})
		return
	}
	if isExist {
		ctx.JSON(200, gin.H{
			"status": "posted",
			"name":   "已经存在",
		})
		return
	}
	ctx.JSON(200, gin.H{
		"status": "posted",
		"name":   "注册成功",
	})
	return
}

func AddUser(ctx *gin.Context) {

}
