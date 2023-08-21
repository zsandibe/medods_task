package service

import (
	"task/internal/repository"
	"task/migration"

	"github.com/golang-jwt/jwt"
)

type service struct {
	storage migration.Storage
}

type AuthService interface {
	CreateToken(userGUID string) (*repository.Jwt, error)
	RefreshToken(tokens *repository.Jwt) (*repository.Jwt, error)
}

type AuthClaims struct {
	*jwt.StandardClaims
	User *repository.UserToken
}

// func NewService(storage migration.Storage) AuthService {
// 	return &service{
// 		storage: storage,
// 	}
// }
