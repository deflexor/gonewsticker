package httpHandlers

import (
	"net/http"
	"log"
	"github.com/deflexor/gonewsticker/httpHandlers/httpUtils"
)

func HandleRequest(w http.ResponseWriter, r *http.Request) {
	log.Println("Incoming Request:", r.Method)
	log.Println("Incoming URI:", r.URL.Path)
	switch r.URL.Path {
	case "/":
		Index(w, r)
		break
	case "/comment":
		if r.Method == http.MethodPost {
			AddComment(w, r)
		} else {
			httpUtils.HandleError(&w, 405, "Method not allowed", "Method not allowed", nil)
		}
		break
	case "/news":
		List(w, r)
		break
	case "/login":
		SteamLogin(w, r)
		break
	case "/login_info":
		LoginInfo(w, r)
		break
	default:
		httpUtils.HandleError(&w, 400, "Bad request", "Bad request", nil)
		break
	}
}
