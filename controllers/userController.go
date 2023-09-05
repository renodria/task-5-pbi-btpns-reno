package controllers

import (
	"net/http"
	"strconv"
	"vix-btpns/app"
	"vix-btpns/database"
	"vix-btpns/helpers"
	"vix-btpns/models"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	var newUser *models.User

	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	if !helpers.Required(newUser.Username) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Username is required"})
		return
	}

	if !helpers.IsEmail(newUser.Email) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid email format"})
		return
	}

	if !helpers.Required(newUser.Password) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Password is required"})
		return
	}

	if !helpers.MinlengthPassword(newUser.Password) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Password must be min 6 characters and max 15"})
		return
	}

	var existingUser models.User
	result := database.DB.Where("email = ?", newUser.Email).First(&existingUser)
	if result.RowsAffected > 0 {
		c.JSON(http.StatusConflict, gin.H{"error": "Email already exists"})
		return
	}

	hashPassword, err := helpers.HashPassword(newUser.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	newUser.Password = hashPassword

	if err := database.DB.Create(&newUser).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Create user successful"})
}

func Login(c *gin.Context) {
	var loginData *app.Login

	if err := c.ShouldBindJSON(&loginData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if !helpers.Required(loginData.Email) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Data incomplete"})
		return
	}

	if !helpers.Required(loginData.Password) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Data incomplete"})
		return
	}

	if !helpers.IsEmail(loginData.Email) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid email format"})
		return
	}

	var existingUser models.User
	result := database.DB.Where("email = ?", loginData.Email).First(&existingUser)
	if result.RowsAffected == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	if !helpers.CheckPassword(loginData.Password, existingUser.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	token, err := helpers.GenerateToken(int(existingUser.ID), existingUser.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", token, 3600*24*30, "", "", false, true)

	c.JSON(http.StatusOK, gin.H{"message": "Login succesfully"})
}

func UpdateUser(c *gin.Context) {
	userId := c.Param("userId")
	userLoginID := c.GetInt("userID")

	intUserId, err := strconv.Atoi(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	if intUserId != userLoginID {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorizitaion"})
		return
	}

	var userData models.User
	result := database.DB.Where("id = ?", userId).First(&userData)
	if result.RowsAffected == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
		return
	}

	var updateData *app.Update
	if err := c.ShouldBindJSON(&updateData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var existingEmail models.User
	res := database.DB.Where("email = ?", updateData.Email).First(&existingEmail)
	if res.RowsAffected > 0 {
		c.JSON(http.StatusConflict, gin.H{"error": "Email already exists"})
		return
	}

	if updateData.Username == " " {
		updateData.Username = userData.Username
	}

	if updateData.Email == " " {
		updateData.Email = userData.Email
	}

	userData.Username = updateData.Username
	userData.Email = updateData.Email
	if err := database.DB.Save(&userData).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Update user successful"})
}

func DeleteUser(c *gin.Context) {
	userId := c.Param("userId")
	userLoginID := c.GetInt("userID")

	intUserId, err := strconv.Atoi(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	if intUserId != userLoginID {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorizitaion"})
		return
	}

	var userData models.User
	result := database.DB.Where("id = ?", userId).First(&userData)
	if result.RowsAffected == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
		return
	}

	if err := database.DB.Delete(&userData).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Delete user successful"})
}
