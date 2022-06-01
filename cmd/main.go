package main

import (
	"context"
	"net/http"
	"time"

	"github.com/badfan/inno-taxi-driver-service/app"
	"github.com/badfan/inno-taxi-driver-service/app/api"
	v1 "github.com/badfan/inno-taxi-driver-service/app/api/v1"
	"github.com/badfan/inno-taxi-driver-service/app/handlers"
	"github.com/badfan/inno-taxi-driver-service/app/resources"
	"github.com/badfan/inno-taxi-driver-service/app/services/auth"
	"github.com/badfan/inno-taxi-driver-service/app/services/driver"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type Server struct {
	httpServer *http.Server
}

func (s *Server) Run(handler http.Handler, port string) error {
	s.httpServer = &http.Server{
		Addr:           ":" + port,
		Handler:        handler,
		MaxHeaderBytes: 1 << 20, // 1 MB
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}

	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}

func InitLogger() *zap.SugaredLogger {
	logger, _ := zap.NewProduction()
	sugarLogger := logger.Sugar()
	return sugarLogger
}

func InitRouter(handler *handlers.Handler) *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()

	apiV1 := v1.NewApiV1(handler)
	apiGroup := api.NewApiGroup(handler, apiV1)
	apiGroup.InitRouterGroups(router)

	return router
}

func main() {
	logger := InitLogger()
	defer logger.Sync()

	apiConfig, err := app.NewAPIConfig()
	if err != nil {
		logger.Errorf("apiconfig error: %s", err.Error())
	}
	dbConfig, err := app.NewDBConfig()
	if err != nil {
		logger.Errorf("dbconfig error: %s", err.Error())
	}

	resource, err := resources.NewResource(dbConfig, logger)
	if err != nil {
		logger.Fatalf("db error: %s", err.Error())
	}
	defer resource.Db.Close()

	authService := auth.NewAuthenticationService(resource, logger)
	driverService := driver.NewDriverService(resource, apiConfig, logger)
	handler := handlers.NewHandler(authService, driverService, logger)

	router := InitRouter(handler)

	server := new(Server)
	if err := server.Run(router, apiConfig.APIPort); err != nil {
		logger.Fatalf("error occured while running http server: %s", err.Error())
	}
}
