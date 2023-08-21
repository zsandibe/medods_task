package service

import (
	"fmt"
	"os"
	"task/internal/repository"

	"github.com/golang-jwt/jwt"
)

func (s *service) CheckAccessToken(accessToken string) (*repository.UserToken, bool) {
	token, err := jwt.ParseWithClaims(accessToken, &AuthClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("ACCESS_KEY")), nil
	})
	if err != nil {
		v, _ := err.(*jwt.ValidationError)
		if v.Errors == jwt.ValidationErrorExpired {
			fmt.Println("time expired")
			return nil, true
		}
		return nil, false
	}

	if claims, ok := token.Claims.(*AuthClaims); ok && token.Valid {
		return claims.User, true
	}
	return nil, false
}
