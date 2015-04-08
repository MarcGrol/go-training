package packaging

import (
	"fmt"
	"github.com/gin-gonic/gin"
    "net/http"
)

func ListenForWebRequests() {
	engine := gin.Default() // HL

	engine.GET("/hello", func(c *gin.Context) {
		c.String(http.StatusOK, fmt.Sprintf("Message from planet %s", "Mars")) // HL
	})

	// will block
	engine.Run(":8080")
}
