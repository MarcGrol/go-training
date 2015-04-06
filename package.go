package mylibrary

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func Doit() {
	engine := gin.Default() // HL

	router.GET("/hello", func(c *gin.Context) {
		c.String(http.StatusOK, fmt.Sprintf("Message from planet %s", "Mars")) // HL
	})

	engine.Run(":8080")
}
