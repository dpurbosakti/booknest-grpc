package api

import (
	"github.com/dpurbosakti/booknest-grpc/internal/view/home"
	"github.com/labstack/echo/v4"
)

func (server *Server) home(c echo.Context) error {
	return render(c, home.Show())
}
