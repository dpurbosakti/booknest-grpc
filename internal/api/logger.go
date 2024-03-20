package api

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
