package main

import (
	"log"
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

func newOAuth2Server() *server.Server {
	manager := manage.NewDefaultManager()
	// token memory store
	manager.MustTokenStorage(store.NewFileTokenStore("./tmp/storage/token.db"))
	manager.MapAccessGenerate(generates.NewJWTAccessGenerate([]byte("00000000"), jwt.SigningMethodHS512))

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
	return server.NewDefaultServer(manager)
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
	srv := newOAuth2Server()
	setupOAuth2ServerConfigcation(srv)

	e := echo.New()
	e.Use(middleware.Recover())
	e.POST("/oauth2/token", func(c echo.Context) error {
		err := srv.HandleTokenRequest(c)
		return err
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
