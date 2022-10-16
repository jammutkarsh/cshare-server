package auth

import (
	"errors"
	"github.com/JammUtkarsh/cshare-server/models"
	"github.com/JammUtkarsh/cshare-server/utils"
	"github.com/JammUtkarsh/cypherDecipher"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
	"os"
	"time"
)

func getJwtKey() []byte {
	utils.LoadEnv(".env")
	return []byte(os.Getenv("JWT_SECRET"))

}

func HashPassword(user models.Users) error {
	db := models.CreateConnection()
	models.CloseConnection(db)
	originalPassword := cypherDecipher.DecipherPassword(user.Password, user.PCount, user.SPCount)
	bytes, err := bcrypt.GenerateFromPassword([]byte(originalPassword), 14)
	if err != nil {
		return err
	}
	password := string(bytes)
	models.InsertPasswordHash(db, user.Username, password)
	return nil
}
func CheckPassword(user models.Users) error {
	db := models.CreateConnection()
	models.CloseConnection(db)
	providedPassword := cypherDecipher.DecipherPassword(user.Password, user.PCount, user.SPCount)
	originalPassword := models.GetPasswordHash(db, user.Username)
	err := bcrypt.CompareHashAndPassword([]byte(originalPassword), []byte(providedPassword))
	if err != nil {
		return err
	}
	return nil
}

func GenerateJWT(username string) (tokenString string, err error) {
	expirationTime := time.Now().Add(365 * time.Hour)
	claims := &models.JWTClaim{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err = token.SignedString(getJwtKey())
	return
}
func ValidateToken(signedToken string) (err error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&models.JWTClaim{},
		func(token *jwt.Token) (interface{}, error) {
			return getJwtKey(), nil
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
