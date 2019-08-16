<h1 align="center">📡 ginprom</h1>
<p align="center">
    <em>Prometheus metrics exporter for Gin.</em>
</p>

### 🔰 Installation

```shell
$ go get -u github.com/chenjiandongx/ginprom
```

### 📝 Usage

It's easy to get started with ginprom, only a few lines of code needed.

```golang
import (
	"github.com/chenjiandongx/ginprom"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
    r := gin.Default()
    // use prometheus metrics exporter middleware.
	r.Use(ginprom.PromMiddleware())

    // register the `/metrices` route.
	r.GET("/metrics", ginprom.PromHandler(promhttp.Handler()))

    // your working routes
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "home"})
    })
}
```

### 📊 Grafana

Although Promethues offers a simple dashboard, Grafana is clearly a better choice. [Grafana configuration](./ginprom-service.json).

![](https://user-images.githubusercontent.com/19553554/63159844-f01e2400-c04e-11e9-8b49-69ff3c3159cb.png)
![](https://user-images.githubusercontent.com/19553554/63159842-eeecf700-c04e-11e9-8f6f-ad0d9dec89ad.png)


### 📃 LICENSE

MIT [©chenjiandongx](https://github.com/chenjiandongx)
