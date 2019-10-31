package controller

import (
	"context"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	_oauth2 "github.com/tamanyan/oauth2-server/app/oauth2"
	"github.com/tamanyan/oauth2-server/oauth2"
	"github.com/tamanyan/oauth2-server/server"
)

// OAuth2Handler  represent the httphandler for oauth2
type OAuth2Handler struct {
	OAuth2Usecase _oauth2.Usecase
}

// NewOAuth2Handler will initialize the articles/ resources endpoint
func NewOAuth2Handler(e *echo.Echo, srv *server.Server, manager oauth2.Manager, us _oauth2.Usecase) {
	handler := &OAuth2Handler{
		OAuth2Usecase: us,
	}
	// e.GET("/articles", handler.FetchArticle)
	// e.POST("/articles", handler.Store)
	// e.GET("/articles/:id", handler.GetByID)
	// e.DELETE("/articles/:id", handler.Delete)
	e.POST("/oauth2/token", handler.IssueAccessToken)

	r := e.Group("/oauth2")
	r.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey:    []byte(os.Getenv("JWT_SECRET_KEY")),
		SigningMethod: "HS512",
	}))
	r.DELETE("/token", func(c echo.Context) error {
		user := c.Get("user").(*jwt.Token)
		err := manager.RemoveAccessToken(user.Raw)

		if err != nil {
			return c.NoContent(http.StatusInternalServerError)
		}

		return c.NoContent(http.StatusOK)
	})
	r.GET("/verify", func(c echo.Context) error {
		ti, err := srv.ValidateBearerToken(c)

		if err != nil {
			return err
		}

		return c.JSON(http.StatusOK, echo.Map{
			"scope":      ti.GetScope(),
			"client_id":  ti.GetClientID(),
			"expires_in": int64(ti.GetAccessExpiresIn() / time.Second),
		})
	})
}

// IssueAccessToken will create access token
func (h *OAuth2Handler) IssueAccessToken(c echo.Context) error {
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	h.OAuth2Usecase.IssueAccessToken(ctx)

	return c.JSON(http.StatusOK, echo.Map{
		"hello": "world",
	})
	// form := form.OAuth2TokenForm{}

	// if err := c.Bind(&form); err != nil {
	// 	log.Println(err)
	// 	return c.String(http.StatusBadRequest, err.Error())
	// }

	// validate := validator.New()
	// err := validate.Struct(form)

	// if err != nil {
	// 	log.Println(err)
	// 	return c.String(http.StatusBadRequest, err.Error())
	// }

	// err = srv.HandleTokenRequest(c)

	// return err
}
