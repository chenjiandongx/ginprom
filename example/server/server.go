package main

import (
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/chenjiandongx/ginprom"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func init() {
	rand.Seed(time.Now().Unix())
}

func zzZ() {
	time.Sleep(time.Millisecond * time.Duration(rand.Int()%1000))
}

func main() {
	r := gin.Default()
	r.Use(ginprom.PromMiddleware())

	r.GET("/metrics", ginprom.PromHandler(promhttp.Handler()))

	r.GET("/", func(c *gin.Context) {
		zzZ()
		c.JSON(http.StatusOK, gin.H{
			"message": "home",
		})
	})

	r.GET("/index", func(c *gin.Context) {
		zzZ()
		c.JSON(http.StatusOK, gin.H{
			"message": "index",
		})
	})

	r.GET("/forbidden", func(c *gin.Context) {
		zzZ()
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
			"message": "forbidden",
		})
	})

	r.GET("/badreq", func(c *gin.Context) {
		zzZ()
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "badreq",
		})
	})

	if err := r.Run(":4433"); err != nil {
		log.Fatalf("run server error: %v", err)
	}
}
