package api

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"golang.org/x/oauth2"
	"golang.org/x/sync/errgroup"

	"github.com/dpurbosakti/booknest-grpc/internal/config"
	db "github.com/dpurbosakti/booknest-grpc/internal/db/sqlc"
	"github.com/getsentry/sentry-go"
	sentryecho "github.com/getsentry/sentry-go/echo"
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

func configSentry() {
	// To initialize Sentry's handler, you need to initialize Sentry itself beforehand
	if err := sentry.Init(sentry.ClientOptions{
		Dsn: "https://334e57cc8ce5f326c1027e3a715f5849@o4506986005987328.ingest.us.sentry.io/4506986019618816",
		// Set TracesSampleRate to 1.0 to capture 100%
		// of transactions for performance monitoring.
		// We recommend adjusting this value in production,
		TracesSampleRate: 1.0,
	}); err != nil {
		fmt.Printf("Sentry initialization failed: %v", err)
	}
}

func (server *Server) setupRouter() {
	router := echo.New()

	// router.Use(LoggerMiddleware)
	router.Use(middleware.Recover())

	configSentry()
	// Open the log file for writing
	file, err := os.OpenFile("logfile.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to open log file")
	}
	defer file.Close()

	// Create a zerolog file writer
	fileLogger := zerolog.New(file).With().Timestamp().Logger()
	router.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Output: fileLogger,
	}))
	router.Logger = &SentryLogger{
		FileLogger: fileLogger,
	}
	router.Use(sentryecho.New(sentryecho.Options{}))
	sentry.CaptureMessage("It works")
	router.Static("/assets", "internal/assets")
	router.GET("/ping", server.ping)
	router.GET("/home", server.home)
	router.GET("/v1/verify_email", server.verify_email)

	auth := router.Group("/auth")
	auth.GET("/v1/google/login", server.googleLogin)
	auth.GET("/google/callback", server.googleCallback)

	server.router = router
}

func (server *Server) Start(address string) error {
	return server.router.Start(address)
}

func (server *Server) Shutdown(ctx context.Context) error {
	return server.router.Shutdown(ctx)
}

type ResponseRecorder struct {
	http.ResponseWriter
	StatusCode int
	Body       []byte
}

func (rec *ResponseRecorder) WriteHeader(statusCode int) {
	rec.StatusCode = statusCode
	rec.ResponseWriter.WriteHeader(statusCode)
}

func (rec *ResponseRecorder) Write(body []byte) (int, error) {
	rec.Body = body
	return rec.ResponseWriter.Write(body)
}

// SentryLogger is a custom logger that logs to both a file and Sentry.
type SentryLogger struct {
	FileLogger zerolog.Logger
}

// Println prints a log message to both file and Sentry.
func (l *SentryLogger) Println(args ...interface{}) {
	l.FileLogger.Print(args...)
	if len(args) > 0 {
		sentry.CaptureMessage(fmt.Sprintf("%v", args[0]))
	}
}
