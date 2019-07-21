package context

import (
	"staff-mall-center/pkg/e"

	"github.com/gin-gonic/gin"
)

type Context struct {
	*gin.Context
}

type HandlerFunc func(c *Context)

func Handle(h HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := &Context{
			c,
		}
		h(ctx)
	}
}

func (ctx *Context) GenResError(code int, msg string) {
	httpcode := code
	if httpcode != 500 {
		httpcode = 200
	}
	ctx.JSON(httpcode, gin.H{
		"errorCode": code,
		"errorMsg":  e.GetMsg(code),
		"info":      msg,
	})
}

func (ctx *Context) GenResSuccess(data interface{}) {
	ctx.JSON(200, gin.H{
		"errorCode": 0,
		"msg":       "success",
		"data":      data,
	})
}
