package controller

import (
	"net/http"
	"strconv"

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

	}
	//create a new record in the database with the data contained in the BlogPost object
	if err := database.DB.Create(&blogPost).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid payload"})

	}
	// Send success response
	c.JSON(http.StatusOK, gin.H{
		"message": "Post created successfully!",
	})

}

// Get All post
func AllPost(c *gin.Context) {
	pageStr := c.Query("page")
	if pageStr == "" {
		pageStr = "1" // Default value
	}
	page, _ := strconv.Atoi(pageStr)
	limit := 5
	offset := (page - 1) * limit
	var total int64
	var getblog []models.Blog
	database.DB.Preload("User").Offset(offset).Limit(limit).Find(&getblog)
	database.DB.Model(&models.Blog{}).Count(&total)
	c.JSON(200, gin.H{
		"data":     getblog,
		"page":     page,
		"pagesize": limit,
		"count":    total,
	})

}

// Details about the post
func DetailPost(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	var blogpost models.Blog
	database.DB.Where("id=?", id).Preload("User").First(&blogpost)
	c.JSON(http.StatusOK, gin.H{
		"data": blogpost,
	})
}
