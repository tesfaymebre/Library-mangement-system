package data

import (
	"context"
	"errors"
	"strings"
	"task_manager/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

// New User registration
func RegisterUser(newUser models.SignUpVerification) error {
	// convert the user email to lowercase
	newUser.Email = strings.ToLower(newUser.Email)

	// check if the user already exists
	var existingUser models.User
	err := userCollection.FindOne(context.TODO(), bson.M{"email": newUser.Email}).Decode(&existingUser)
	if err != nil && err != mongo.ErrNoDocuments {
		// If there's an error other than no documents found, return the error
		return err
	} else if err == nil {
		// If user exists, return user already exists error
		return errors.New("user already exists")
	}

	// user registration logic using bcrypt to hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	newUser.Password = string(hashedPassword)

	// check if the user is the signupCollection, if so update the user details
	var existingSignup models.SignUpVerification
	err = signupCollection.FindOne(context.TODO(), bson.M{"email": newUser.Email}).Decode(&existingSignup)
	if err != nil && err != mongo.ErrNoDocuments {
		// If there's an error other than no documents found, return the error
		return err
	} else if err == nil {
		// If a signup entry exists, delete it
		_, err = signupCollection.DeleteOne(context.TODO(), bson.M{"email": newUser.Email})
		if err != nil {
			return err
		}
	}

	// insert the user details into the signupCollection
	_, err = signupCollection.InsertOne(context.TODO(), newUser)
	if err != nil {
		return err
	}

	return nil
}

// otp verification
func VerifyOTP(otpVerification models.OtpVerification) (models.User, error) {
	// convert the user email to lowercase
	otpVerification.Email = strings.ToLower(otpVerification.Email)

	// check if the user exists in the signupCollection
	var existingUser models.SignUpVerification
	err := signupCollection.FindOne(context.TODO(), bson.M{"email": otpVerification.Email}).Decode(&existingUser)
	if err != nil {
		return models.User{}, errors.New("user not found")
	}

	// check if the OTP is valid
	if otpVerification.Otp != existingUser.GenratedOtp {
		return models.User{}, errors.New("invalid OTP")
	}

	// check if the OTP has expired
	if existingUser.OtpExpiryTime.Before(time.Now()) {
		return models.User{}, errors.New("OTP has expired")
	}

	// parse the user details to the user struct
	var user models.User
	user.Email = existingUser.Email
	user.Password = existingUser.Password

	// set user role by counting the existing users
	count, err := userCollection.CountDocuments(context.TODO(), bson.M{})
	if err != nil {
		return models.User{}, err
	}

	if count == 0 {
		user.Role = "admin"
	} else {
		user.Role = "user"
	}

	// insert the user details into the userCollection
	_, err = userCollection.InsertOne(context.TODO(), user)
	if err != nil {
		return models.User{}, err
	}

	// delete the user details from the signupCollection
	_, err = signupCollection.DeleteOne(context.TODO(), bson.M{"email": otpVerification.Email})
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

// resend OTP
func ResendOTP(resendOTP models.ResendOTPRequest) error {
	// convert the user email to lowercase
	resendOTP.Email = strings.ToLower(resendOTP.Email)

	// check if the user exists in the signupCollection
	var existingUser models.SignUpVerification
	err := signupCollection.FindOne(context.TODO(), bson.M{"email": resendOTP.Email}).Decode(&existingUser)
	if err != nil {
		return errors.New("user not found")
	}

	// update the OTP and expiry time in the signupCollection
	_, err = signupCollection.UpdateOne(context.TODO(), bson.M{"email": resendOTP.Email}, bson.M{"$set": bson.M{"genrated_otp": resendOTP.Otp, "otp_expiry_time": resendOTP.OtpExpiryTime}})
	if err != nil {
		return err
	}

	return nil
}

// User login
func Login(user models.User) (models.User, error) {
	// convert the user email to lowercase
	user.Email = strings.ToLower(user.Email)

	var existingUser models.User
	err := userCollection.FindOne(context.TODO(), bson.M{"email": user.Email}).Decode(&existingUser)

	if err != nil || bcrypt.CompareHashAndPassword([]byte(existingUser.Password), []byte(user.Password)) != nil {
		return models.User{}, errors.New("invalid Email or Password")
	}

	return existingUser, nil
}

// Get all users
func GetAllUsers() ([]models.User, error) {
	var users []models.User

	cursor, err := userCollection.Find(context.Background(), bson.M{})
	if err != nil {
		return users, err
	}

	if err = cursor.All(context.Background(), &users); err != nil {
		return users, err
	}

	return users, nil
}

// get user by id
func GetUserByID(id string) (models.User, error) {
	var user models.User
	err := userCollection.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&user)
	if err != nil {
		return models.User{}, errors.New("user not found")
	}

	return user, nil
}

// promote user to admin
func PromoteUser(id string) error {
	// check if the user exists
	var user models.User
	err := userCollection.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&user)
	if err != nil {
		return errors.New("user not found")
	}

	// update the user role to admin
	_, err = userCollection.UpdateOne(context.TODO(), bson.M{"_id": id}, bson.M{"$set": bson.M{"role": "admin"}})
	if err != nil {
		return err
	}

	return nil
}
