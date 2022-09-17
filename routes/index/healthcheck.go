package index

import "github.com/gin-gonic/gin";
import (
	"time"
)

func HealthcheckRoutes(route *gin.Engine){
	healthcheck := route.Group("/")
	healthcheck.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "OK",
			"timestamp": time.Now().Unix(),
		})
	})
}
