package http

import (
	"log"
	"net/http"
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/tamanyan/oauth2-server/oauth2"
	"github.com/tamanyan/oauth2-server/server"
)

// NewOAuth2TokenHandler will initialize the articles/ resources endpoint
func NewOAuth2TokenHandler(e *echo.Echo, srv *server.Server, manager oauth2.Manager) {
	// handler := &ArticleHandler{
	// 	AUsecase: us,
	// }
	// e.GET("/articles", handler.FetchArticle)
	// e.POST("/articles", handler.Store)
	// e.GET("/articles/:id", handler.GetByID)
	// e.DELETE("/articles/:id", handler.Delete)
	e.POST("/oauth2/token", func(c echo.Context) error {
		err := srv.HandleTokenRequest(c)
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

		log.Println(user.Raw)
		if err != nil {
			return c.NoContent(http.StatusInternalServerError)
		}

		return c.NoContent(http.StatusOK)
	})
}
