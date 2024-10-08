package main

import (
	"fmt"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// Global variable to store users in-memory
var users = make(map[string]*User)

// global variable to store jwt secret key
var jwtKey = []byte("your_jwt_secret")

// middleware to check the JWT token
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(401, gin.H{"error": "Authorization header is required"})
			c.Abort()
			return
		}

		authParts := strings.Split(authHeader, " ")
		if len(authParts) != 2 || authParts[0] != "Bearer" {
			c.JSON(401, gin.H{"error": "Authorization header format must be Bearer {token}"})
			c.Abort()
			return
		}

		// Parse JWT token
		token, err := jwt.Parse(authParts[1], func(token *jwt.Token) (interface{}, error) {
			// Don't forget to validate the alg is what you expect:
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}

			return jwtKey, nil
		})

		if err != nil || !token.Valid {
			c.JSON(401, gin.H{"error": "Invalid JWT"})
			c.Abort()
			return
		}

		c.Next()
	}
}

func main() {

	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to the Go Authentication and Authorization tutorial!",
		})
	})

	// registration

	router.POST("/register", func(c *gin.Context) {
		var user User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(400, gin.H{"error": "Invalid request payload"})
			return
		}

		// User registration logic using bcrypt to hash the password
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err != nil {
			c.JSON(500, gin.H{"error": "Failed to hash the password"})
			return
		}

		user.Password = string(hashedPassword)
		users[user.Email] = &user

		c.JSON(200, gin.H{"message": "User registered successfully"})
	})

	// login
	router.POST("/login", func(c *gin.Context) {
		var user User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(400, gin.H{"error": "Invalid request payload"})
			return
		}

		// Check if the user exists
		storedUser, found := users[user.Email]
		if !found || bcrypt.CompareHashAndPassword([]byte(storedUser.Password), []byte(user.Password)) != nil {
			c.JSON(401, gin.H{"error": "Invalid Email or Password"})
			return
		}

		// Generate JWT token
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"user_id": storedUser.ID,
			"email":   storedUser.Email,
		})

		jwtToken, err := token.SignedString(jwtKey)

		if err != nil {
			c.JSON(500, gin.H{"error": "Failed to generate JWT token"})
			return
		}

		c.JSON(200, gin.H{"message": "User logged in successfully", "token": jwtToken})
	})

	router.GET("/secure", AuthMiddleware(), func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "You are authorized to access this secure endpoint!",
		})
	})

	router.Run()
}
