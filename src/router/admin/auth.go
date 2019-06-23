package admin

import "github.com/gin-gonic/gin"

func AdminLogin(c *gin.Context) {
	resp := map[string]string{"hello": "world"}

	c.JSON(200, resp)
}
