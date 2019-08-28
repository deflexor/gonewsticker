package httpHandlers

import (
	"log"
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
	pd := PageData{"Новости", "{}"}
	ses, err := r.Cookie("SESSIONV2")
	if err == nil {
		id := ses.Value
		userData, err1 := session.GetUser(id)
		if err1 == nil {
			user1, err2 := json.Marshal(userData)
			if err2 == nil {
				pd.User = string(user1)
			} else {
				log.Printf("json.Marshal(userData) err: %s\n", err2)
			}
		} else {
			log.Printf("session.GetUser err: %s\n", err1)
		}
	}
	renderTemplate(w, &pd)
}
