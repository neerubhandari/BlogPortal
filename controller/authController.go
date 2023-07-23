package controller

import (
	"fmt"
	"net/http"
	"regexp"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/neerubhandari/BlogPortal/database"
	"github.com/neerubhandari/BlogPortal/models"
	"github.com/neerubhandari/BlogPortal/utils"
	"gorm.io/gorm"
)

func validateEmail(email string) bool {
	Re := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	return Re.MatchString(email)
}

func Register(c *gin.Context) {

	var data map[string]interface{}
	var user models.User

	// Bind the request body to data
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Unable to parse request body"})
		return
	}

	// Validate the password length
	if len(data["password"].(string)) < 6 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Password must be at least 6 characters long"})
		return
	}

	// Validate the email format
	email := strings.TrimSpace(data["email"].(string))
	if !validateEmail(email) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid email address"})
		return
	}

	// Check if email already exist in the database
	if err := database.DB.Where("email = ?", email).First(&user).Error; err != gorm.ErrRecordNotFound {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email already exists"})
		return
	}

	// Prepare the new user record
	user = models.User{
		FirstName: data["first_name"].(string),
		LastName:  data["last_name"].(string),
		Phone:     data["phone"].(string),
		Email:     data["email"].(string),
	}

	// Set the password (assuming SetPassword() hashes the password)
	user.SetPassword(data["password"].(string))

	// Create the new user
	if err := database.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	// Send success response
	c.JSON(http.StatusOK, gin.H{
		"message": "Account created successfully",
		"user":    user,
	})
}

func Login(c *gin.Context) {
	var data map[string]string

	if err := c.ShouldBindJSON(&data); err != nil {
		fmt.Println("Unable to parse body")
	}
	var user models.User
	database.DB.Where("email =?", data["email"]).First(&user)
	if user.Id == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "Email address doesn't exist"})
		return
	}

	if err := user.ComparePassword(data["password"]); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Incorrect password"})
		return
	}
	token, err := utils.GenerateJwt(strconv.Itoa(int(user.Id)))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// Set a new cookie
	c.SetCookie(
		"jwt",    // Cookie name
		token,    // Cookie value
		48*60*60, // Max age in seconds
		"/",      // Path
		"",       // Domain, empty for current domain
		false,    // Secure only
		true,     // HttpOnly
	)
	c.JSON(http.StatusOK, gin.H{"message": "You have successfully login"})
	return

}
