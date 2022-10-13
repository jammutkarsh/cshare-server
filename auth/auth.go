package auth

import (
	"errors"
	"github.com/JammUtkarsh/cshare-server/models"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
	"time"
)

var jwtKey = []byte("SecretYouShouldHide")

func HashPassword(user models.Users) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	if err != nil {
		return err
	}
	password := string(bytes)
	models.InsertPasswordHash(user.Username, password)
	return nil
}
func CheckPassword(username, providedPassword string) error {
	originalPassword := models.GetPasswordHash(username)
	err := bcrypt.CompareHashAndPassword([]byte(originalPassword), []byte(providedPassword))
	if err != nil {
		return err
	}
	return nil
}

func GenerateJWT(username string) (tokenString string, err error) {
	expirationTime := time.Now().Add(24 * 7 * 4 * time.Hour)
	claims := &models.JWTClaim{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err = token.SignedString(jwtKey)
	return
}
func ValidateToken(signedToken string) (err error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&models.JWTClaim{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtKey), nil
		},
	)
	if err != nil {
		return
	}
	claims, ok := token.Claims.(*models.JWTClaim)
	if !ok {
		err = errors.New("couldn't parse claims")
		return
	}
	if claims.ExpiresAt < time.Now().Local().Unix() {
		err = errors.New("token expired")
		return
	}
	return
}
