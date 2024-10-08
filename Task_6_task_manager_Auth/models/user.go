package models

import "time"

type User struct {
	ID       uint   `bson:"id"`
	Email    string `bson:"email"`
	Password string `bson:"password"`
	Role     string `bson:"role"`
}

// signup verification
type SignUpVerification struct {
	Email         string    `bson:"email"`
	Password      string    `bson:"password"`
	GenratedOtp   string    `bson:"genrated_otp"`
	OtpExpiryTime time.Time `bson:"otp_expiry_time"`
}

// otp verification
type OtpVerification struct {
	Email string `bson:"email"`
	Otp   string `bson:"otp"`
}

// resend otp request
type ResendOTPRequest struct {
	Email         string    `bson:"email"`
	Otp           string    `bson:"otp"`
	OtpExpiryTime time.Time `bson:"otp_expiry_time"`
}

// login response
type LoginResponse struct {
	AccessToken  string `bson:"access_token"`
	RefreshToken string `bson:"refresh_token"`
	ID           uint   `bson:"id"`
	Email        string `bson:"email"`
	Role         string `bson:"role"`
}
