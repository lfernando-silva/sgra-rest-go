package index

import "github.com/gin-gonic/gin";

import (
	"github.com/lfernando-silva/sgra-rest-go/controllers/index"
)

func SessionRoutes(route *gin.Engine){
	session := route.Group("/")
	session.POST("/signup", controllers.Signup)
}
