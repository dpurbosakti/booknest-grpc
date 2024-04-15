package api

import (
	"net/http"

	"github.com/dpurbosakti/booknest-grpc/internal/view/http_error"
	"github.com/labstack/echo/v4"
)

func customHTTPErrorHandler(err error, c echo.Context) {
	if he, ok := err.(*echo.HTTPError); ok {
		if he.Code == http.StatusNotFound {
			// c.String(http.StatusNotFound, "404 Not Found")
			errorEndPointNotFound(c)
			return
		}
	}
	// Handle other errors
	c.String(http.StatusInternalServerError, "Internal Server Error")
}

func errorEndPointNotFound(c echo.Context) error {
	return render(c, http_error.Show())
}
