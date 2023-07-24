package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/neerubhandari/BlogPortal/database"
	"github.com/neerubhandari/BlogPortal/models"
)

// create post
func CreatePost(c *gin.Context) {
	var blogPost models.Blog
	// Bind the request body to data
	if err := c.ShouldBindJSON(&blogPost); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Unable to parse request body"})
		return
	}
	//create a new record in the database with the data contained in the BlogPost object
	if err := database.DB.Create(&blogPost).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid payload"})
		return
	}
	// Send success response
	c.JSON(http.StatusOK, gin.H{
		"message": "Post created successfully!",
	})

}
