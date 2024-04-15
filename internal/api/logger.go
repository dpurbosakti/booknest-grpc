package api

import (
	"net/http"
	"os"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// import (
// 	"net/http"

// 	"github.com/rs/zerolog/log"

// 	"github.com/labstack/echo/v4"
// )

// func makeLogger(c echo.Context) {
// 	logger := log.Info()

// 	// if c.Request().Response.StatusCode != http.StatusOK {
// 	// 	body, err := io.ReadAll(c.Request().Body)
// 	// 	if err != nil {
// 	// 		log.Error().Err(err)
// 	// 	}
// 	// 	logger = log.Error().Bytes("body", body)
// 	// }

// 	logger.
// 		Str("protocol", "http").
// 		Str("method", c.Request().Method).
// 		Str("path", c.Path()).
// 		Int("status_code", c.Request().Response.StatusCode).
// 		Str("status_text", http.StatusText(c.Request().Response.StatusCode)).
// 		// Dur("duration", c.Request().).
// 		Msg("received a HTTP request")
// }

// func logMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
// 	return func(c echo.Context) error {
// 		makeLogger(c)
// 		return next(c)
// 	}
// }

func LoggerMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		start := time.Now()

		if err := next(c); err != nil {
			c.Error(err)
		}

		end := time.Now()
		latency := end.Sub(start)

		logger := log.Info()
		status := c.Response().Status
		if status != http.StatusOK {
			logger = log.Error()
		}
		// Log request details using Zerolog
		logger.
			Str("method", c.Request().Method).
			Str("uri", c.Request().RequestURI).
			Int("status", status).
			Dur("latency", latency).
			Msg("request handled")

		return nil
	}
}

func setupLogger() {
	runLogFile, _ := os.OpenFile(
		"myapp.log",
		os.O_APPEND|os.O_CREATE|os.O_WRONLY,
		0664,
	)
	multi := zerolog.MultiLevelWriter(os.Stdout, runLogFile)
	log.Logger = zerolog.New(multi).With().Timestamp().Logger()

	log.Info().Msg("Hello World!")

}
