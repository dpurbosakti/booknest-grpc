package api

import (
	"context"
	"errors"
	"net/http"

	"github.com/rs/zerolog/log"
	"golang.org/x/oauth2"
	"golang.org/x/sync/errgroup"

	"github.com/dpurbosakti/booknest-grpc/internal/config"
	db "github.com/dpurbosakti/booknest-grpc/internal/db/sqlc"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Server struct {
	config    config.Config
	store     db.Store
	router    *echo.Echo
	googleCfg *oauth2.Config
}

func NewServer(config config.Config, store db.Store) (*Server, error) {
	server := &Server{
		config:    config,
		store:     store,
		googleCfg: getGoogleCfg(config),
	}

	server.setupRouter()
	return server, nil
}

func RunEchoServer(ctx context.Context, waitGroup *errgroup.Group, config config.Config, store db.Store) {
	server, err := NewServer(config, store)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot create server")
	}

	waitGroup.Go(func() error {
		log.Info().Msgf("start echo server at %s", config.EchoServerAddress)
		err = server.Start(config.EchoServerAddress)
		if err != nil {
			if errors.Is(err, http.ErrServerClosed) {
				return nil
			}
			log.Info().Err(err).Msg("cannot start server")
			return err
		}

		return nil
	})

	waitGroup.Go(func() error {
		<-ctx.Done()
		log.Info().Msg("graceful shudown echo server")

		err := server.Shutdown(context.Background())
		if err != nil {
			log.Error().Err(err).Msg("failed to shutdown echo server")
			return err
		}

		log.Info().Msg("echo server is stopped")
		return nil
	})
}

func (server *Server) setupRouter() {
	router := echo.New()
	logger := log.Info()
	router.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogURI:    true,
		LogStatus: true,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			logger.
				Str("URI", v.URI).
				Int("status", v.Status).
				Msg("request")

			return nil
		},
	}))
	router.GET("/health-check", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "ok")
	})

	auth := router.Group("/auth")
	auth.GET("/v1/google/login", server.GoogleLogin)
	auth.GET("/google/callback", server.GoogleCallback)

	server.router = router
}

func (server *Server) Start(address string) error {
	return server.router.Start(address)
}

func (server *Server) Shutdown(ctx context.Context) error {
	return server.router.Shutdown(ctx)
}
