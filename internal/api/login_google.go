package api

import (
	"encoding/json"
	"net/http"

	"github.com/dpurbosakti/booknest-grpc/internal/config"
	"github.com/labstack/echo/v4"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type GoogleResponse struct {
	Email            string `json:"email"`
	Name             string `json:"name"`
	VerifiedEmail    bool   `json:"verified_email"`
	OauthAccessToken string `json:"oauth_access_token"`
}

func (server *Server) googleLogin(c echo.Context) error {

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

func (server *Server) googleCallback(c echo.Context) error {

	// state
	state := c.QueryParam("state")
	if state != "" {
		c.JSON(http.StatusUnprocessableEntity, map[string]any{"message": "state does not match"})
	}

	// code
	code := c.QueryParam("code")

	// exchange code for token
	token, err := server.googleCfg.Exchange(c.Request().Context(), code)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, map[string]any{"message": err.Error()})
	}
	if token == nil {
		return c.JSON(http.StatusInternalServerError, map[string]any{"message": "no token"})
	} else {
		// use google api to get user info
		resp, err := http.Get(server.config.GoogleTokenAccessURL + token.AccessToken)

		if err != nil {
			return c.JSON(http.StatusUnprocessableEntity, map[string]any{"message": err.Error()})
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			return c.JSON(resp.StatusCode, map[string]any{"message": "Non-OK status code received"})
		}

		// Parse the response body as JSON
		var userData *GoogleResponse
		err = json.NewDecoder(resp.Body).Decode(&userData)
		if err != nil {
			return c.JSON(http.StatusUnprocessableEntity, map[string]any{"message": err.Error()})
		}
		return c.JSON(http.StatusOK, map[string]any{"data": userData})
	}

}
