package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
);
import (
	"github.com/lfernando-silva/sgra-rest-go/database/models"
	"github.com/lfernando-silva/sgra-rest-go/database"
)

type SignupInput struct {
	Name string `json:"name" binding:"required"`
	Document string `json:"document" binding:"required"`
	Email string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
  }

func Signup(c *gin.Context){
	db := database.GetDB()
  var input SignupInput
  if err := c.ShouldBindJSON(&input); err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    return
  }

  // Create book
  user := models.User{
	Name: input.Name, 
	Document: input.Document,
	Email: input.Document,
	Password: input.Password,
  }

  db.Create(&user)
  
  c.JSON(http.StatusOK, gin.H{"data": user})
}