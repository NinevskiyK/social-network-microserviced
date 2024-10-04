package utils

import (
	"errors"
	"fmt"
	"os"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

var Now = time.Now

func GetIdFromJWT(tokenString string) (id string, err error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(os.Getenv("hmacSecret")), nil
	})
	if err != nil {
		return "", err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		exp_time, is_string := claims["exp_time"].(string)
		if !is_string {
			return "", errors.New("invalid token")
		}
		token_time, err := time.Parse(time.RFC3339, exp_time)
		if err != nil {
			return "", err
		}
		if token_time.Compare(Now()) == -1 {
			return "", errors.New("expired token")
		}
		id, is_string := claims["id"].(string)
		if !is_string {
			return "", errors.New("invalid token")
		}
		return id, nil
	}
	return "", errors.New("invalid token")
}

func CreateJWT(id string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":       id,
		"exp_time": Now().Add(60 * time.Minute).Format(time.RFC3339),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("hmacSecret")))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
