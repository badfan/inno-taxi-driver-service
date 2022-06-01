package driver

import (
	"context"
	"crypto/sha1"
	"errors"
	"fmt"
	"time"

	"github.com/badfan/inno-taxi-driver-service/app"
	"github.com/badfan/inno-taxi-driver-service/app/models"
	"github.com/badfan/inno-taxi-driver-service/app/resources"
	"github.com/dgrijalva/jwt-go" //nolint:typecheck
	"go.uber.org/zap"
)

const (
	hashSalt   = "384dds87jf98l!^#1hub"
	SigningKey = "3y74hhf3740&#!f493!5"
)

type IDriverService interface {
	SignUp(ctx context.Context, driver *models.Driver) (int, error)
	SignIn(ctx context.Context, phone, password string) (string, error)
	GetDriverRating(ctx context.Context, id int) (float32, error)
	GetDriverStatus(ctx context.Context, id int) (bool, error)
}

type DriverService struct {
	resource  resources.IResource
	apiConfig *app.APIConfig
	logger    *zap.SugaredLogger
}

type TokenClaims struct {
	jwt.StandardClaims       //nolint:typecheck
	ID                 int32 `json:"id"`
}

func NewDriverService(resource resources.IResource, apiConfig *app.APIConfig, logger *zap.SugaredLogger) *DriverService {
	return &DriverService{resource: resource, apiConfig: apiConfig, logger: logger}
}

func (s *DriverService) SignUp(ctx context.Context, driver *models.Driver) (int, error) {
	if _, err := s.resource.GetDriverIDByPhone(ctx, driver.PhoneNumber); err == nil {
		return 0, errors.New("phone number is already taken")
	}

	driver.Password = generatePasswordHash(driver.Password)

	res, err := s.resource.CreateDriver(ctx, driver)
	if err != nil {
		return 0, err
	}

	return res, nil
}

func (s *DriverService) SignIn(ctx context.Context, phone, password string) (string, error) {
	password = generatePasswordHash(password)

	driver, err := s.resource.GetDriverByPhoneAndPassword(ctx, phone, password)
	if err != nil {
		return "", errors.New("invalid phone number or password")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &TokenClaims{ //nolint:typecheck
		StandardClaims: jwt.StandardClaims{ //nolint:typecheck
			ExpiresAt: time.Now().Add(time.Duration(s.apiConfig.TokenTTL) * time.Hour).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		ID: driver.ID,
	})

	return token.SignedString([]byte(SigningKey))
}

func (s *DriverService) GetDriverRating(ctx context.Context, id int) (float32, error) {
	return s.resource.GetDriverRatingByID(ctx, id)
}

func (s *DriverService) GetDriverStatus(ctx context.Context, id int) (bool, error) {
	return s.resource.GetDriverStatusByID(ctx, id)
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(hashSalt)))
}
