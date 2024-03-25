package api

import (
	"github.com/dpurbosakti/booknest-grpc/internal/view/ping"
	"github.com/labstack/echo/v4"
)

func (server *Server) ping(c echo.Context) error {
	return render(c, ping.Show())
}
