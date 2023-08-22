package service

import (
	"encoding/base64"
	"log"
	"os"
	"task/internal/repository"
	"task/pkg/hash"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

func (s *service) CreateAccessToken(userGUID string) (string, error) {
	accessTokenClaims := &jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Hour * 1).Unix(),
		Id:        userGUID,
		IssuedAt:  time.Now().Add(time.Second * 5).Unix(),
		Subject:   "user",
	}

	setNewClaims := jwt.NewWithClaims(jwt.SigningMethodHS512, accessTokenClaims)
	token, err := setNewClaims.SignedString([]byte(os.Getenv("ACCESS_KEY")))
	if err != nil {
		return "", err
	}
	return token, nil
}

func (s *service) CreateRefreshToken() string {
	refreshToken := uuid.New()
	encodedRefreshToken := base64.StdEncoding.EncodeToString([]byte(refreshToken.String()))
	return encodedRefreshToken
}

func (s *service) CreatePair(userGUID string) (*repository.Jwt, *repository.UserToken, error) {
	accessToken, err := s.CreateAccessToken(userGUID)
	if err != nil {
		log.Println("error creating access token : %v\n", err)
		return nil, nil, err
	}
	refreshToken := s.CreateRefreshToken()
	hashedRefreshToken, err := hash.Generate(refreshToken)
	if err != nil {
		log.Println("error generating refresh token : %v\n", err)
		return nil, nil, err
	}

	bindTokens := accessToken[len(accessToken)-6:]

	tokensClient := &repository.Jwt{
		UserGUID:     userGUID,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}
	tokensDB := &repository.UserToken{
		UserGUID:     userGUID,
		RefreshToken: hashedRefreshToken,
		BindTokens:   bindTokens,
	}

	return tokensClient, tokensDB, nil
}

func (s *service) CreateToken(userGUID string) (*repository.Jwt, error) {
	tokensClient, tokensDB, err := s.CreatePair(userGUID)
	// fmt.Println("tC", tokensClient, "\n", "tDB", tokensDB)
	if err != nil {
		log.Println("error creating pair : %v\n", err)
		return nil, err
	}
	s.storage.Create(tokensDB)
	// fmt.Println(tokensClient)
	return tokensClient, nil
}
