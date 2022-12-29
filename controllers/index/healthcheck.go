package controllers

import (
	"github.com/gin-gonic/gin"
	"time"
);

func HealthCheck(ctx *gin.Context){
	ctx.JSON(200, gin.H{
		"message": "OK",
		"timestamp": time.Now().Unix(),
	})
}