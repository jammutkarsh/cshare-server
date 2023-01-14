package auth

// auth package deals with authentication and authorization.
// It has methods related to
// 1. Hashing and verifying passwords
// 2. JWT creation & validation

import (
	"errors"
	"log"
	"os"
	"time"

	"github.com/JammUtkarsh/cshare-server/models"
	"github.com/JammUtkarsh/cshare-server/utils"
	"github.com/JammUtkarsh/cypherDecipher"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

type jwtClaim struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

// getJwtKey reads JWT_SECRET key from .env file and returns the secret in bytes.
func getJwtKey() []byte {
	utils.LoadEnv(".env")
	return []byte(os.Getenv("JWT_SECRET"))

}

// HashPassword removes the salt from the provided password returns the hashed version of it.
// Internally it uses a special library github.com/JammUtkarsh/cypherDecipher, to remove salt from the password.
func HashPassword(user models.Users) (password string, err error) {
	var bytes []byte
	originalPassword := cypherDecipher.DecipherPassword(user.Password, user.PCount, user.SPCount)
	bytes, err = bcrypt.GenerateFromPassword([]byte(originalPassword), 14)
	if err != nil {
		return "", nil
	}
	return string(bytes), nil
}

// CheckPassword removes salt from the provided password and verifies its hash.
// Internally it uses a special library github.com/JammUtkarsh/cypherDecipher, to remove salt from the password.
func CheckPassword(user models.Users) error {
	db, err := models.CreateConnection()
	defer models.CloseConnection(db)
	if err != nil {
		return err
	}

	var originalPassword string

	if originalPassword, err = models.GetPasswordHash(db, user.Username); err != nil {
		return err
	}

	providedPassword := cypherDecipher.DecipherPassword(user.Password, user.PCount, user.SPCount)
	if err := bcrypt.CompareHashAndPassword([]byte(originalPassword), []byte(providedPassword)); err != nil {
		return err
	}
	return nil
}

// GenerateJWT takes username as parameter and returns a JWT string and error if any.
func GenerateJWT(username string) (tokenString string, err error) {
	utils.LoadEnv(".env")
	timeFactor, err := time.ParseDuration(os.Getenv("TIME_FACTOR"))
	if err != nil {
		timeFactor = 24
		log.Println("TIME_FACTOR is invlaid or not set. Defaulting to 24 hours.")
	}
	expirationTime := time.Now().Add(timeFactor * time.Hour)

	claims := &jwtClaim{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err = token.SignedString(getJwtKey())
	return
}

// ValidateToken takes JWT string as a parameter and verifies it. Returns an error if invalid.
func ValidateToken(signedToken string) (err error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&jwtClaim{},
		func(token *jwt.Token) (interface{}, error) {
			return getJwtKey(), nil
		},
	)
	if err != nil {
		return
	}

	claims, ok := token.Claims.(*jwtClaim)
	if !ok {
		err = errors.New("couldn't_parse_claims")
		return
	}

	if claims.ExpiresAt < time.Now().Local().Unix() {
		err = errors.New("token_expired")
		return
	}
	return
}
