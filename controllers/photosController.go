package controllers

import (
	"net/http"
	"os"
	"time"
	"vix-btpns/database"
	"vix-btpns/models"

	"github.com/gin-gonic/gin"
)

func ShowPhotos(c *gin.Context) {
	userId := c.GetInt("userID")

	var photos []models.Photos
	result := database.DB.Where("user_id = ?", userId).Find(&photos)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch photos"})
		return
	}

	if len(photos) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Photos not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": photos})
}

func ShowPhotoId(c *gin.Context) {
	userId := c.GetInt("userID")
	photoId := c.Param("photoId")

	var photo models.Photos
	result := database.DB.Where("user_id = ? AND id = ?", userId, photoId).First(&photo)
	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Photo not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": photo})
}

func CreatePhoto(c *gin.Context) {
	userLoginID := c.GetInt("userID")
	newPhoto := &models.Photos{}

	title := c.PostForm("title")
	caption := c.PostForm("caption")

	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	now := time.Now().Format("20060102150405")
	newFilename := now + "_" + file.Filename

	err = c.SaveUploadedFile(file, "./uploads/"+newFilename)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file"})
		return
	}

	newPhoto.Title = title
	newPhoto.Caption = caption
	newPhoto.PhotoUrl = "http://localhost:" + os.Getenv("PORT") + "/uploads/" + newFilename
	newPhoto.UserID = userLoginID

	if err := database.DB.Create(&newPhoto).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create photo"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Create photo successfully"})
}

func UpdatePhoto(c *gin.Context) {
	userId := c.GetInt("userID")
	photoId := c.Param("photoId")

	var findPhoto models.Photos
	result := database.DB.Where("user_id = ? AND id = ?", userId, photoId).First(&findPhoto)
	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Photo not found"})
		return
	}

	title := c.PostForm("title")
	caption := c.PostForm("caption")

	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	now := time.Now().Format("20060102150405")
	newFilename := now + "_" + file.Filename

	err = c.SaveUploadedFile(file, "./uploads/"+newFilename)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file"})
		return
	}

	findPhoto.Title = title
	findPhoto.Caption = caption
	findPhoto.PhotoUrl = "http://localhost:" + os.Getenv("PORT") + "/uploads/" + newFilename

	if err := database.DB.Save(&findPhoto).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update photo"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Updated photo successfully"})
}

func DeletePhoto(c *gin.Context) {
	userId := c.GetInt("userID")
	photoId := c.Param("photoId")

	var findPhoto models.Photos
	result := database.DB.Where("user_id = ? AND id = ?", userId, photoId).First(&findPhoto)
	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Photo not found"})
		return
	}

	if err := database.DB.Delete(&findPhoto).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete photo"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Delete photo successfully"})
}
