package index

import "github.com/gin-gonic/gin";
import (
	"github.com/lfernando-silva/sgra-rest-go/controllers/index"
)

func HealthcheckRoutes(route *gin.Engine){
	healthcheck := route.Group("/")
	healthcheck.GET("/", controllers.HealthCheck)
}
