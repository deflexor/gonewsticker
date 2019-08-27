package httpHandlers

import (
	"github.com/deflexor/gonewsticker/session"
	"github.com/deflexor/gonewsticker/httpHandlers/httpUtils"
	"net/http"
)

func LoginInfo(w http.ResponseWriter, r *http.Request) {
	ses, err := r.Cookie("SESSIONV2")
	if err != nil {
		httpUtils.HandleError(&w, 400, "Session cookie not found!", "", nil)
		return
	}
	id := ses.Value
	user, err := session.GetUser(id)
	if err != nil {
		httpUtils.HandleError(&w, 400, err.Error(), "", nil)
		return		
	}
	httpUtils.HandleSuccess(&w, user)
}
