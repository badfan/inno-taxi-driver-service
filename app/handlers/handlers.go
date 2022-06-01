package handlers

import (
	"github.com/badfan/inno-taxi-driver-service/app/services/auth"
	"github.com/badfan/inno-taxi-driver-service/app/services/driver"
	"go.uber.org/zap"
)

type Handler struct {
	authService   auth.IAuthenticationService
	driverService driver.IDriverService
	logger        *zap.SugaredLogger
}

func NewHandler(authService auth.IAuthenticationService, driverService driver.IDriverService, logger *zap.SugaredLogger) *Handler {
	return &Handler{authService: authService, driverService: driverService, logger: logger}
}
