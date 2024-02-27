package lesson_4

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func Gin() {
	r := gin.Default()

	r.GET("/time", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"time_from_gin": time.Now(),
		})
	})

	r.Run(":8080")
}
