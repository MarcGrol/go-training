package mylibrary

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func ListenForWebRequests() {
	engine := gin.Default() // HL

	router.GET("/hello", func(c *gin.Context) {
		c.String(http.StatusOK, fmt.Sprintf("Message from planet %s", "Mars")) // HL
	})

	// will block
	engine.Run(":8080")
}
