package ginprom

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func init() {
	rand.Seed(time.Now().Unix())
}

func zzZ() {
	time.Sleep(time.Millisecond * time.Duration(rand.Int()%1250))
}

func floodRequest() {
	// reused client object
	client := &http.Client{}
	endpoints := []string{"/", "/index", "/forbidden", "/badreq"}
	for {
		u := fmt.Sprintf("http://localhost:4433%s", endpoints[rand.Int()%4])
		req, _ := http.NewRequest(http.MethodGet, u, nil)
		if _, err := client.Do(req); err != nil {
			log.Printf("request error: %v", err)
			// something wrong, zzZ...
		}
		zzZ()
	}
}

func main() {
	r := gin.Default()
	r.Use(PromMiddleware())

	r.GET("/metrics", PromHandler(promhttp.Handler()))

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

	go floodRequest()

	if err := r.Run(":4433"); err != nil {
		log.Fatalf("run server error: %v", err)
	}
}
