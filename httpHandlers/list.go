package httpHandlers

import (
	"github.com/deflexor/gonewsticker/storage"
	"github.com/deflexor/gonewsticker/httpHandlers/httpUtils"
	"net/http"

//	"github.com/deflexor/gonewsticker/httpHandlers/httpUtils"
)

func List(w http.ResponseWriter, r *http.Request) {
	httpUtils.HandleSuccess(&w, storage.Get())
}
