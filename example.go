package ginprom

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	r := gin.Default()
	r.Use(PromMiddleware())

	r.GET("/metrics", PromHandler(promhttp.Handler()))

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "home",
		})
	})

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.GET("/index", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "index",
		})
	})

	r.GET("/forbidden", func(c *gin.Context) {
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
			"message": "forbidden",
		})
	})

	if err := r.Run(":4433"); err != nil {
		log.Fatalf("run server error: %v", err)
	}
}
