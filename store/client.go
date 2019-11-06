package store

import (
	// "errors"
	// "sync"
	"encoding/json"
	// "fmt"
	// "log"

	"github.com/tidwall/buntdb"

	"github.com/tamanyan/oauth2-server/models"
	"github.com/tamanyan/oauth2-server/oauth2"
)

// NewMemoryClientStore create a token store instance based on memory
func NewMemoryClientStore() (client *ClientStore, err error) {
	client, err = NewClientStore(":memory:")
	return
}

// NewClientStore create client store
func NewClientStore(filename string) (client *ClientStore, err error) {
	db, err := buntdb.Open(filename)
	if err != nil {
		return
	}
	client = &ClientStore{
		db: db,
	}
	return
}

// ClientStore client information store
type ClientStore struct {
	db *buntdb.DB
}

// GetByID according to the ID for the client information
func (cs *ClientStore) GetByID(id string) (cli oauth2.ClientInfo, err error) {
	verr := cs.db.View(func(tx *buntdb.Tx) (err error) {
		jv, err := tx.Get(id)
		if err != nil {
			return
		}
		var cm models.Client
		err = json.Unmarshal([]byte(jv), &cm)
		if err != nil {
			return
		}
		cli = &cm
		return
	})
	if verr != nil {
		if verr == buntdb.ErrNotFound {
			return
		}
		err = verr
	}
	return
}

// Set set client information
func (cs *ClientStore) Set(id string, cli oauth2.ClientInfo) (err error) {
	jv, err := json.Marshal(cli)
	if err != nil {
		return
	}
	err = cs.db.Update(func(tx *buntdb.Tx) error {
		_, _, err := tx.Set(id, string(jv), nil)
		return err
	})
	return
}
