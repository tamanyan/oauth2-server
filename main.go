package main

import (
	"log"
	"net/http"
	// "fmt"

	// goerrors "github.com/pkg/errors"
	"github.com/tamanyan/oauth2-server/oauth2"
	"github.com/tamanyan/oauth2-server/errors"
	"github.com/tamanyan/oauth2-server/manage"
	"github.com/tamanyan/oauth2-server/models"
	"github.com/tamanyan/oauth2-server/server"
	"github.com/tamanyan/oauth2-server/store"
	// "golang.org/x/oauth2/clientcredentials"
)

const (
	authServerURL = "http://localhost:9096"
)

func main() {
	manager := manage.NewDefaultManager()
	// token memory store
	manager.MustTokenStorage(store.NewFileTokenStore("./tmp/storage/token.db"))

	// client memory store
	clientStore, err := store.NewClientStore("./tmp/storage/client.db")
	if err != nil {
		log.Fatal(err)
	}
	clientStore.Set("000000", &models.Client{
		ID:     "000000",
		Secret: "999999",
		Domain: "http://localhost",
	})
	manager.MapClientStorage(clientStore)

	srv := server.NewDefaultServer(manager)
	srv.SetAllowGetAccessRequest(true)
	srv.SetClientInfoHandler(server.ClientFormHandler)

	srv.SetClientAuthorizedHandler(func(clientID string, grant oauth2.GrantType) (allowed bool, err error) {
		log.Println(clientID, grant)
		allowed = true
		return
	})

	srv.SetInternalErrorHandler(func(err error) (re *errors.Response) {
		log.Println("Internal Error:", err.Error())
		// gerr := goerrors.New("error")
		// gerr = goerrors.Wrap(err, "open failed")
		// gerr = goerrors.Wrap(err, "read config failed")

		// fmt.Printf("%+v\n", gerr)
		return
	})

	srv.SetResponseErrorHandler(func(re *errors.Response) {
		log.Println("Response Error:", re.Error.Error())
	})

	http.HandleFunc("/authorize", func(w http.ResponseWriter, r *http.Request) {
		err := srv.HandleAuthorizeRequest(w, r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	})

	srv.SetPasswordAuthorizationHandler(func(username, password string) (userID string, err error) {
		log.Println(username, password)
		if username == "test" && password == "test" {
			userID = "test"
			return
		}
		err = errors.ErrInvalidGrant
		return
	})

	http.HandleFunc("/token", func(w http.ResponseWriter, r *http.Request) {
		err := srv.HandleTokenRequest(w, r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	log.Println("Ready on http://localhost:9096")

	log.Fatal(http.ListenAndServe(":9096", nil))
}
