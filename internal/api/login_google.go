package api

import (
	"net/http"

	"github.com/dpurbosakti/booknest-grpc/internal/config"
	"github.com/labstack/echo/v4"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

func (server *Server) GoogleLogin(c echo.Context) error {

	url := server.googleCfg.AuthCodeURL(server.config.GoogleState)

	// redirect to google login page
	return c.Redirect(http.StatusTemporaryRedirect, url)
}

func getGoogleCfg(config config.Config) *oauth2.Config {
	conf := &oauth2.Config{
		ClientID:     config.GoogleCientID,
		ClientSecret: config.GoogleClientSecret,
		RedirectURL:  config.GoogleRedirectURL,
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email",
			"https://www.googleapis.com/auth/userinfo.profile",
			"https://www.googleapis.com/auth/calendar",
			"https://www.googleapis.com/auth/calendar.events",
		},
		Endpoint: google.Endpoint,
	}

	return conf
}
