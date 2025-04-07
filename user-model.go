package main

import (
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	Email    string `json:"email" gorm:"unique"`
	Password string `json:"password"`
}

func createUser(db *gorm.DB, user *User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user.Password = string(hashedPassword)
	result := db.Create(user)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func login(db *gorm.DB, user *User) (string, error) {
	// get user from email
	selectedUser := new(User)
	result := db.Where("email = ?", user.Email).First(selectedUser)
	if result.Error != nil {
		return "", result.Error
	}

	// compare password
	err := bcrypt.CompareHashAndPassword(
		[]byte(selectedUser.Password), // password ที่มาจาก DB
		[]byte(user.Password))         // password ที่มาจาก request body
	if err != nil {
		return "", err
	}

	// pass = return JWT
	// Create JWT token

	jwtSecretKey := "TestSecret" // shound be env

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = selectedUser.ID
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	t, err := token.SignedString([]byte(jwtSecretKey))
	if err != nil {
		return "", err
	}

	return t, nil
}
