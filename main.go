package main

import (
	"log"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	_ "github.com/joho/godotenv/autoload"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	_oauth2Controller "github.com/tamanyan/oauth2-server/app/oauth2/http/controller"
	_oauth2Usecase "github.com/tamanyan/oauth2-server/app/oauth2/usecase"
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
	manager.MapAccessGenerate(generates.NewJWTAccessGenerate([]byte(os.Getenv("JWT_SECRET_KEY")), jwt.SigningMethodHS512))

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
	// e.Use(middleware.CORS())
	e.Use(middleware.Gzip())
	e.Use(middleware.RequestID())

	timeoutContext := time.Duration(1000 * time.Second)
	au := _oauth2Usecase.NewOAuth2Usecase(timeoutContext)
	_oauth2Controller.NewOAuth2Handler(e, srv, manager, au)

	if os.Getenv("DEBUG") == "1" {
		e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
			Format: "method=${method}, uri=${uri}, status=${status} latency=${latency_human}\n",
		}))
	} else {
		e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
			Format: `{"time":"${time_rfc3339_nano}","id":"${id}","remote_ip":"${remote_ip}","host":"${host}",` +
				`"method":"${method}","uri":"${uri}","user_agent":"${user_agent}",status":${status},"error":"${error}",` +
				`"latency":${latency},"latency_human":"${latency_human}","bytes_in":${bytes_in},` +
				`"bytes_out":${bytes_out}}` + "\n",
		}))
	}

	var port = "9096"

	if len(os.Getenv("PORT")) > 0 {
		port = os.Getenv("PORT")
	}

	e.Logger.Fatal(e.Start(":" + port))
}
