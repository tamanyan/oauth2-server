package main

import (
	"log"
	"net/http"
	"os"

	"github.com/dgrijalva/jwt-go"
	_ "github.com/joho/godotenv/autoload"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/tamanyan/oauth2-server/errors"
	"github.com/tamanyan/oauth2-server/generates"
	"github.com/tamanyan/oauth2-server/manage"
	"github.com/tamanyan/oauth2-server/models"
	"github.com/tamanyan/oauth2-server/oauth2"
	"github.com/tamanyan/oauth2-server/server"
	"github.com/tamanyan/oauth2-server/store"
)

var (
	secretKey = "sample_pwd"
)

func newManager() oauth2.Manager {
	manager := manage.NewDefaultManager()
	// token memory store
	manager.MustTokenStorage(store.NewFileTokenStore("./tmp/storage/token.db"))
	manager.MapAccessGenerate(generates.NewJWTAccessGenerate([]byte(secretKey), jwt.SigningMethodHS512))

	// client memory store
	clientStore, err := store.NewClientStore("./tmp/storage/client.db")
	if err != nil {
		log.Fatal(err)
	}
	clientStore.Set("sample", &models.Client{
		ID:     "sample",
		Secret: "999999",
		Domain: "http://localhost",
	})
	manager.MapClientStorage(clientStore)
	return manager
}

func setupOAuth2ServerConfigcation(srv *server.Server) {
	srv.SetAllowGetAccessRequest(false)
	srv.SetClientInfoHandler(server.ClientFormHandler)

	srv.SetClientAuthorizedHandler(func(clientID string, grant oauth2.GrantType) (allowed bool, err error) {
		// log.Println(clientID, grant)
		allowed = true
		return
	})

	srv.SetInternalErrorHandler(func(err error) (re *errors.Response) {
		log.Println("Internal Error:", err.Error())
		return
	})

	srv.SetResponseErrorHandler(func(re *errors.Response) {
		log.Println("Response Error:", re.Error.Error())
	})

	srv.SetPasswordAuthorizationHandler(func(username, password string) (userID string, err error) {
		// log.Println(username, password)
		if username == "test" && password == "test" {
			userID = "test"
			return
		}
		err = errors.ErrInvalidGrant
		return
	})
}

func main() {
	manager := newManager()
	srv := server.NewDefaultServer(manager)
	setupOAuth2ServerConfigcation(srv)

	e := echo.New()
	e.Use(middleware.Recover())
	e.Use(middleware.BodyLimit("1M"))
	e.Use(middleware.CORS())
	e.Use(middleware.Gzip())
	e.Use(middleware.RequestID())
	e.POST("/oauth2/token", func(c echo.Context) error {
		err := srv.HandleTokenRequest(c)
		return err
	})

	r := e.Group("/oauth2")
	r.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey:    []byte(secretKey),
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

	if os.Getenv("DEBUG") == "1" {
		e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
			Format: "method=${method}, uri=${uri}, status=${status} latency=${latency_human}\n",
		}))
	} else {
		e.Use(middleware.Logger())
	}

	var port = "9096"

	if len(os.Getenv("PORT")) > 0 {
		port = os.Getenv("PORT")
	}

	e.Logger.Fatal(e.Start(":" + port))
}
