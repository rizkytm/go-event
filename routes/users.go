package routes

import (
	"net/http"

	"example.com/go-event/models"
	"example.com/go-event/utils"
	"github.com/gin-gonic/gin"
)

func signup(context *gin.Context) {
	var user models.User

	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse"})
		return
	}

	err = user.Save()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not save"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "user created"})
}

func login(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse"})
		return
	}
	err = user.ValidateCredentials()
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "invalid credentials"})
		return
	}
	token, err := utils.GenerateToken(user.Email, user.ID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not auth"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "login success", "token": token})
}
