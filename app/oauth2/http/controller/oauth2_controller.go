package controller

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"gopkg.in/go-playground/validator.v9"

	form "github.com/tamanyan/oauth2-server/app/oauth2/http/form"
	"github.com/tamanyan/oauth2-server/oauth2"
	"github.com/tamanyan/oauth2-server/server"
)

// NewOAuth2Handler will initialize the articles/ resources endpoint
func NewOAuth2Handler(e *echo.Echo, srv *server.Server, manager oauth2.Manager) {
	// handler := &ArticleHandler{
	// 	AUsecase: us,
	// }
	// e.GET("/articles", handler.FetchArticle)
	// e.POST("/articles", handler.Store)
	// e.GET("/articles/:id", handler.GetByID)
	// e.DELETE("/articles/:id", handler.Delete)
	e.POST("/oauth2/token", func(c echo.Context) error {
		form := form.OAuth2TokenForm{}

		if err := c.Bind(&form); err != nil {
			log.Println(err)
			return c.String(http.StatusBadRequest, err.Error())
		}

		validate := validator.New()
		err := validate.Struct(form)

		if err != nil {
			log.Println(err)
			return c.String(http.StatusBadRequest, err.Error())
		}

		err = srv.HandleTokenRequest(c)

		return err
	})

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
