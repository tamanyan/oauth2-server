package main

import (
	"log"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	_ "github.com/joho/godotenv/autoload"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	_adminController "github.com/tamanyan/oauth2-server/app/admin/http/controller"
	_adminRepository "github.com/tamanyan/oauth2-server/app/admin/repository"
	_adminUsecase "github.com/tamanyan/oauth2-server/app/admin/usecase"
	_middleware "github.com/tamanyan/oauth2-server/app/middleware"
	_oauth2Controller "github.com/tamanyan/oauth2-server/app/oauth2/http/controller"
	_oauth2Usecase "github.com/tamanyan/oauth2-server/app/oauth2/usecase"
	_profileController "github.com/tamanyan/oauth2-server/app/profile/http/controller"
	_profileRepository "github.com/tamanyan/oauth2-server/app/profile/repository"
	_profileUsecase "github.com/tamanyan/oauth2-server/app/profile/usecase"
	"github.com/tamanyan/oauth2"
	"github.com/tamanyan/oauth2/generates"
	"github.com/tamanyan/oauth2/manage"
	"github.com/tamanyan/oauth2/models"
	"github.com/tamanyan/oauth2/store"
)

var (
	secretKey = "sample_pwd"
)

func newManager() oauth2.Manager {
	manager := manage.NewDefaultManager()
	// token memory store
	manager.MustTokenStorage(store.NewTokenStore(store.NewTokenConfig("./tmp/storage/oauth2.db", "sqlite3", "oauth2_token"), 600))
	manager.MapAccessGenerate(generates.NewJWTAccessGenerate([]byte(os.Getenv("JWT_SECRET_KEY")), jwt.SigningMethodHS512))

	// client memory store
	clientStore, err := store.NewClientStore(store.NewClientConfig("./tmp/storage/oauth2.db", "sqlite3", "client"))
	if err != nil {
		log.Fatal(err)
	}
	err = clientStore.Set("sample", &models.Client{
		ID:     "sample",
		Secret: "999999",
		Domain: "http://localhost",
	})
	if err != nil {
		log.Fatal(err)
	}
	manager.MapClientStorage(clientStore)
	return manager
}

func main() {
	e := echo.New()
	e.Use(middleware.Recover())
	e.Use(middleware.BodyLimit("1M"))
	// e.Use(middleware.CORS())
	e.Use(middleware.Gzip())
	e.Use(middleware.RequestID())

	manager := newManager()
	goMiddleware := _middleware.InitMiddleware()
	timeoutContext := time.Duration(1000 * time.Second)

	au := _oauth2Usecase.NewOAuth2Usecase(manager, timeoutContext)
	_oauth2Controller.NewOAuth2Handler(e, goMiddleware, manager, au)

	pr := _profileRepository.NewProfileRepository()
	pu := _profileUsecase.NewProfileUsecase(pr)
	_profileController.NewProfileHandler(e, goMiddleware, pu)

	adr := _adminRepository.NewAdminRepository()
	adu := _adminUsecase.NewAdminUsecase(adr)
	_adminController.NewAdminHandler(e, goMiddleware, adu)

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
