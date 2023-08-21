package service

import (
	"errors"
	"log"
	"task/internal/repository"
	"task/pkg/hash"
)

func (s *service) UpdateToken(tokens *repository.Jwt) (*repository.Jwt, error) {
	_, valid := s.CheckAccessToken(tokens.AccessToken)

	if !valid {
		log.Printf("Invalid access token: --> %v\n", valid)
		return nil, errors.New("Invalid access token")
	}

	bindTokens := tokens.AccessToken[len(tokens.AccessToken)-6:]

	oldRefreshToken, err := s.storage.Get(tokens.UserGUID, bindTokens)
	if err != nil {
		log.Printf("Didn't find the second pair of token: %v\n", err)
		return nil, err
	}

	isValidRefreshToken := hash.Compare(oldRefreshToken.RefreshToken, tokens.RefreshToken)

	if !isValidRefreshToken {
		log.Println("Invalid refresh token")
		return nil, errors.New("Invalid refresh token")
	}

	newJwtTokens, newRefreshToken, err := s.CreatePair(tokens.UserGUID)
	if err != nil {
		log.Printf("Couldn create the new pair of token: %v\n", err)
		return nil, err
	}
	s.storage.Update(oldRefreshToken, newRefreshToken)
	return newJwtTokens, nil
}
