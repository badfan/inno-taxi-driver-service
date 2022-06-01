package auth

import (
	"errors"

	"github.com/badfan/inno-taxi-driver-service/app/resources"
	"github.com/badfan/inno-taxi-driver-service/app/services/driver"
	"github.com/dgrijalva/jwt-go" //nolint:typecheck
	"go.uber.org/zap"
)

type IAuthenticationService interface {
	ParseToken(accessToken string) (int, error)
}

type AuthenticationService struct {
	resource resources.IResource
	logger   *zap.SugaredLogger
}

func NewAuthenticationService(resource resources.IResource, logger *zap.SugaredLogger) *AuthenticationService {
	return &AuthenticationService{resource: resource, logger: logger}
}
func (s *AuthenticationService) ParseToken(accessToken string) (int, error) {
	token, err := jwt.ParseWithClaims(accessToken, &driver.TokenClaims{}, func(token *jwt.Token) (interface{}, error) { //nolint:typecheck
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte(driver.SigningKey), nil
	})
	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(*driver.TokenClaims)
	if !ok {
		return 0, errors.New("token claims are not of type *tokenClaims")
	}

	return int(claims.ID), nil
}
