package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const secretKey = "supersecret"

func GenerateToken(email string, userId int64) (string, error) {
	/* creates a new [Token] with the specified
	* signing method and claims. */
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":  email,
		"userId": userId,
		"exp":    time.Now().Add(time.Hour * 2).Unix(),
	})
	return token.SignedString([]byte(secretKey))
}

func VerifyToken(token string) (int64, error) {
	// parses, validates, verifies the signature and returns the parsed token.
	// keyFunc will receive the parsed token and should return the cryptographic key for verifying the signature.
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		// check if token was signed with specific way.
		_, ok := token.Method.(*jwt.SigningMethodHMAC) // type checking syntax. returns bool
		if !ok {
			return nil, errors.New("Unexpected signing method.")
		}
		return []byte(secretKey), nil
	})
	if err != nil {
		return 0, errors.New("Could not parse token.")
	}
	tokenIsValid := parsedToken.Valid
	if !tokenIsValid {
		return 0, errors.New("Invalid token.")
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return 0, errors.New("Invalid token claims.")
	}

	// email := claims["email"].(string)
	userId := int64(claims["userId"].(float64))
	return userId, nil
}
