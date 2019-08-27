package httpHandlers

import (
	"net/http"
	"text/template"
	"encoding/json"
	"github.com/deflexor/gonewsticker/session"
)

type PageData struct {
	Title string
	User  string
}

func renderTemplate(w http.ResponseWriter, p *PageData) {
	var t = template.Must(template.New("index.html").Delims("<%", "%>").ParseFiles("./views/index.html"))
	err := t.ExecuteTemplate(w, "index.html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func Index(w http.ResponseWriter, r *http.Request) {
	// http.Redirect(w, r, "/edit/"+title, http.StatusFound)
	ses, err := r.Cookie("SESSIONV2")
	user := "{}"
	if err == nil {
		id := ses.Value
		userData, err1 := session.GetUser(id)
		if err1 == nil {
			user1, err2 := json.Marshal(userData)
			if err2 != nil {
				user = string(user1)
			}
		}
	}
	pd := PageData{"Новости", user}
	renderTemplate(w, &pd)
}
