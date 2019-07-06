package authcheck

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func AuthRequired(c *gin.Context) {

	fmt.Println("Auth Check")

	// Pass on to the next-in-chain

	c.Next()
}
