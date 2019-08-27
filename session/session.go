package session

import (
	"errors"
	"log"
	"sync"
	"time"

	"github.com/deflexor/gonewsticker/structs"
)

type Session struct {
	T time.Time
	S structs.PlayerSummary
}

var store = make(map[string]Session)
var mux sync.Mutex
var maxSessionAge = time.Duration(time.Hour * 8000)

func GetUser(id string) (structs.PlayerSummary, error) {
	user, ok := store[id]
	if !ok {
		return structs.PlayerSummary{}, errors.New("user not found")
	}
	now := time.Now()
	if now.Sub(user.T) > maxSessionAge {
		log.Printf("Session expired for user: %s(%s)\n", id, user.S.PersonaName)
		delete(store, id)
		return structs.PlayerSummary{}, errors.New("user not found")
	}
	return user.S, nil
}

func Login(data structs.PlayerSummary) {
	mux.Lock()
	defer mux.Unlock()

	now := time.Now()
	store[data.SteamId] = Session{now, data}
}
