package controllers

import (
	"errors"
	"net/http"
	"os"
	"strconv"
	"task_manager/data"
	"task_manager/models"
	"task_manager/utils" // Assuming you have a utils package for sending OTPs
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

// access token generation
func createAccessToken(user models.User, expiry int) (accessToken string, err error) {
	// create a new JWT token
	expirationTime := time.Now().Add(time.Duration(expiry) * time.Minute)
	claims := &models.JwtCustomClaims{
		UserID: strconv.FormatUint(uint64(user.ID), 10),
		Email:  user.Email,
		Role:   user.Role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	// Load the secret key from the .env file
	secretKey := os.Getenv("JWT_SECRET")
	if secretKey == "" {
		return "", errors.New("jwt secret key not set in environment variables")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	accessToken, err = token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return accessToken, nil
}

// refresh token generation
func createRefreshToken(user models.User, expiry int) (refreshToken string, err error) {
	expirationTime := time.Now().Add(time.Duration(expiry) * time.Hour)
	claimsRefresh := &models.JwtCustomRefreshClaims{
		ID: strconv.FormatUint(uint64(user.ID), 10),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	// Load the secret key from the .env file
	secretKey := os.Getenv("JWT_SECRET")

	if secretKey == "" {
		return "", errors.New("JWT secret key not set in environment variables")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claimsRefresh)

	refreshToken, err = token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return refreshToken, nil
}

// New user registration endpoint
func RegisterUser(c *gin.Context) {
	var newUser models.SignUpVerification
	if err := c.BindJSON(&newUser); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "invalid request payload"})
		return
	}

	// Generate OTP
	otp, err := utils.GenerateOTP(newUser.Email)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "failed to generate OTP"})
		return
	}

	// save otp in newUser
	newUser.GenratedOtp = otp
	newUser.OtpExpiryTime = time.Now().Add(time.Minute * 5)

	if err := data.RegisterUser(newUser); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "internal server error"})
		return
	}

	// Send OTP
	if err := utils.SendOTP(newUser.Email, otp); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "failed to send OTP"})
		return
	}

	c.IndentedJSON(http.StatusCreated, gin.H{"message": "user registered successfully, please check your email to verify your account"})
}

// Verify OTP endpoint
func VerifyOTP(c *gin.Context) {
	var otpVerification models.OtpVerification
	if err := c.BindJSON(&otpVerification); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "invalid request payload"})
		return
	}

	user, err := data.VerifyOTP(otpVerification)

	if err != nil {
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
		return
	}

	// generate access token
	expiryMinutes, err := strconv.Atoi(os.Getenv("ACCESS_TOKEN_EXPIRY_MINUTES"))
	if err != nil {
		expiryMinutes = 15
	}

	accessToken, err := createAccessToken(user, expiryMinutes)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "failed to generate access token"})
		return
	}

	// generate refresh token
	expiryHours, err := strconv.Atoi(os.Getenv("REFRESH_TOKEN_EXPIRY_HOURS"))
	if err != nil {
		expiryHours = 24
	}
	refreshToken, err := createRefreshToken(user, expiryHours)

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "failed to generate refresh token"})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{
		"message":       "OTP verified successfully",
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	})
}

// resend OTP endpoint
func ResendOTP(c *gin.Context) {
	var resendOTP models.ResendOTPRequest
	if err := c.BindJSON(&resendOTP); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "invalid request payload"})
		return
	}

	// Generate OTP
	otp, err := utils.GenerateOTP(resendOTP.Email)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "failed to generate OTP"})
		return
	}

	// save otp in resendOTP
	resendOTP.Otp = otp
	resendOTP.OtpExpiryTime = time.Now().Add(time.Minute * 5)

	if err := data.ResendOTP(resendOTP); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "internal server error"})
		return
	}

	// Send OTP
	if err := utils.SendOTP(resendOTP.Email, otp); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "failed to send OTP"})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "OTP sent successfully"})
}

// User login endpoint
func Login(c *gin.Context) {
	var user models.User
	if err := c.BindJSON(&user); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "invalid request payload"})
		return
	}

	existingUserInfo, err := data.Login(user)
	if err != nil {
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
		return
	}

	// generate access token
	expiryMinutes, err := strconv.Atoi(os.Getenv("ACCESS_TOKEN_EXPIRY_MINUTES"))
	if err != nil {
		expiryMinutes = 15
	}

	access_token, err := createAccessToken(existingUserInfo, expiryMinutes)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "failed to generate access token"})
		return
	}

	// generate refresh token
	expiryHours, err := strconv.Atoi(os.Getenv("REFRESH_TOKEN_EXPIRY_HOURS"))
	if err != nil {
		expiryHours = 24
	}

	refresh_token, err := createRefreshToken(existingUserInfo, expiryHours)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "failed to generate refresh token"})
		return
	}

	// login response
	loginResponse := models.LoginResponse{
		AccessToken:  access_token,
		RefreshToken: refresh_token,
		ID:           existingUserInfo.ID,
		Email:        existingUserInfo.Email,
		Role:         existingUserInfo.Role,
	}

	c.IndentedJSON(http.StatusOK, loginResponse)
}

// Refresh token endpoint
func RefreshToken(c *gin.Context) {
	var refreshToken models.RefreshToken
	if err := c.BindJSON(&refreshToken); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "invalid request payload"})
		return
	}

	// Load the secret key from the .env file
	secretKey := os.Getenv("JWT_SECRET")
	if secretKey == "" {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "JWT secret key not set in environment variables"})
		return
	}

	claims := &models.JwtCustomRefreshClaims{}
	token, err := jwt.ParseWithClaims(refreshToken.RefreshToken, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})

	if err != nil || !token.Valid {
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"message": "Invalid token"})
		return
	}

	userID, err := strconv.ParseUint(claims.ID, 10, 64)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "failed to parse user ID"})
		return
	}

	user, err := data.GetUserByID(strconv.FormatUint(userID, 10))
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "user not found"})
		return
	}

	// generate access token
	expiryMinutes, err := strconv.Atoi(os.Getenv("ACCESS_TOKEN_EXPIRY_MINUTES"))
	if err != nil {
		expiryMinutes = 15
	}

	accessToken, err := createAccessToken(user, expiryMinutes)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "failed to generate access token"})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"access_token": accessToken})
}

// Get all users endpoint
func GetAllUsers(c *gin.Context) {
	users, err := data.GetAllUsers()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "internal server error"})
		return
	}

	c.IndentedJSON(http.StatusOK, users)
}

// Promote user to admin endpoint
func PromoteUser(c *gin.Context) {
	userID := c.Param("id")
	if userID == "" {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "user ID is required"})
		return
	}

	err := data.PromoteUser(userID)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "user promoted to admin successfully"})
}
