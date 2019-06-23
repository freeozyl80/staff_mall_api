package benchmark

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func MyBenchLogger(c *gin.Context) {
	fmt.Println("Im a benchmark!")

	c.Next()
}
