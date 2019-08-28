package storage

import (
	"time"
	"sync"
	"github.com/deflexor/gonewsticker/structs"
)

var mux sync.Mutex
var store structs.NewsList

func Get() structs.NewsList {
	return store
}

func Add(message structs.NewsMessage) {
	mux.Lock()
	defer mux.Unlock()
	store = append(store, message)
}

func AddMany(messages []structs.NewsMessage) int {
	mux.Lock()
	defer mux.Unlock()
	if len(store) == 0 {
		store = messages
	} else {
		cutoffTime := store[0].Created.Add(time.Second)
		for _, m := range messages {
			if m.Created.Before(cutoffTime)  {
				break
			} else {
				store = append([]structs.NewsMessage{m}, store...)
			}
		}
		// store = append(store, messages...)
	}
	return 0
}

func AddComment(guid string, c structs.Comment) bool {
	mux.Lock()
	defer mux.Unlock()

	for i, _ := range store {
		if store[i].GUID == guid {
			c.Added = time.Now()
			store[i].Comments = append(store[i].Comments, c)
			return true
		}
	}
	return false
}

func Clear() {
	store = nil
}
