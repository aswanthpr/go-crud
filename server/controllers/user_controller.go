package controllers

import (
	"crud-app/configs"
	"crud-app/dtos"
	"crud-app/models"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func UserSignup(c *gin.Context) {
	var form dtos.UserSignUpFormDTO

	if err := c.ShouldBindJSON(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "invalid request body",
			"details": err.Error(),
		})
		return
	}

	var existing models.User

	if err := configs.DB.Where("email = ?", form.Email).First(&existing).Error; err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to check existing user",
		})
		return
	} else if err == nil {
		c.JSON(http.StatusConflict, gin.H{
			"error": "email is already registered",
		})
		return
	}

	
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(form.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to process password",
		})
		return
	}

	user := models.User{
		Name:     form.Name,
		Email:    form.Email,
		Password: string(hashedPassword),
	}

	if err := configs.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to create user",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"success": true, "message": "Singup Successful"})

}

func LoginUser(c *gin.Context) {
	var form dtos.UserLoginFormDTO

	if err := c.ShouldBindJSON(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	var user models.User
	err := configs.DB.Where("email = ?", form.Email).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid email or password"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "database error"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(form.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid email or password"})
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "message": "Login success"})
}
