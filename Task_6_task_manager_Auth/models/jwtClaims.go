package models

import "github.com/golang-jwt/jwt"

// JWTClaims represents the structure of the JWT claims
type JwtCustomClaims struct {
	UserID string `json:"user_id"`
	Email  string `json:"email"`
	Role   string `json:"role"`
	jwt.StandardClaims
}

type JwtCustomRefreshClaims struct {
	ID string `json:"id"`
	jwt.StandardClaims
}

type RefreshToken struct {
	RefreshToken string `json:"refresh_token"`
}
