package api

import (
	"errors"
	"net/http"
	"strconv"

	db "github.com/dpurbosakti/booknest-grpc/internal/db/sqlc"
	"github.com/dpurbosakti/booknest-grpc/internal/view/verify_email"
	"github.com/labstack/echo/v4"
)

func (server *Server) verify_email(c echo.Context) error {
	email_id, _ := strconv.Atoi(c.QueryParam("email_id"))

	txResult, err := server.store.VerifyEmailTx(c.Request().Context(), db.VerifyEmailTxParams{
		EmailId:    int64(email_id),
		SecretCode: c.QueryParam("secret_code"),
	})
	if err != nil {
		if errors.Is(err, db.ErrRecordNotFound) {
			return c.JSON(http.StatusNotFound, "failed to verify email: user with email_id %d not found")
		}
		return c.JSON(http.StatusInternalServerError, "failed to verify email: %s")
	}
	return render(c, verify_email.Show(txResult.User.Name))
}
