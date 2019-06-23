package admin

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func AuthRequired(c *gin.Context) {

	fmt.Println("Im a admin!")

	c.Next()
}
