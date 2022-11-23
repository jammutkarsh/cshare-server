package auth

// auth package deals with authentication and authorization related methods.
// It has methods related JWT creation and validation & hashing and validation of hashed passwords.

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

type JWTClaim struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

// getJwtKey reads JWT_SECRET key from .env file and returns the secret
func getJwtKey() []byte {
	utils.LoadEnv(".env")
	return []byte(os.Getenv("JWT_SECRET"))

}

// HashPassword hashes and stores the password in the database.
// Internally it uses a special library github.com/JammUtkarsh/cypherDecipher, to remove salt from the password.
func HashPassword(user models.Users) (password string, err error) {
	db := models.CreateConnection()
	defer models.CloseConnection(db)

	originalPassword := cypherDecipher.DecipherPassword(user.Password, user.PCount, user.SPCount)

	bytes, err := bcrypt.GenerateFromPassword([]byte(originalPassword), 14)
	if err != nil {
		return "", nil
	}
	return string(bytes), nil
}

// CheckPassword takes the user's credentials, removes salt from the password and verifies for the password.
// Internally it uses a special library github.com/JammUtkarsh/cypherDecipher, to remove salt from the password.
func CheckPassword(user models.Users) error {
	db := models.CreateConnection()
	defer models.CloseConnection(db)

	providedPassword := cypherDecipher.DecipherPassword(user.Password, user.PCount, user.SPCount)
	_, originalPassword := models.GetPasswordHash(db, user.Username)

	if err := bcrypt.CompareHashAndPassword([]byte(originalPassword), []byte(providedPassword)); err != nil {
		return err
	}
	return nil
}

// GenerateJWT takes username as a parameter and returns JWT and an error if any.
func GenerateJWT(username string) (tokenString string, err error) {
	expirationTime := time.Now().Add(365 * time.Hour)

	claims := &JWTClaim{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err = token.SignedString(getJwtKey())
	return
}

// ValidateToken takes JWT as a parameter and validates it. Returns error if any
func ValidateToken(signedToken string) (err error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&JWTClaim{},
		func(token *jwt.Token) (interface{}, error) {
			return getJwtKey(), nil
		},
	)
	if err != nil {
		return
	}

	claims, ok := token.Claims.(*JWTClaim)
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
