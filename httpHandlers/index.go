package httpHandlers

import (
	"net/http"
	"text/template"
//	"github.com/deflexor/gonewsticker/httpHandlers/httpUtils"
)

type PageData struct {
	title string
}

var t = template.Must(template.New("index.html").Delims("<%", "%>").ParseFiles("../views/index.html"))

func renderTemplate(w http.ResponseWriter, p *PageData) {
	err := t.ExecuteTemplate(w, "index.html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func Index(w http.ResponseWriter, r *http.Request) {
	// http.Redirect(w, r, "/edit/"+title, http.StatusFound)
	pd := PageData { "Новости" }
	renderTemplate(w, &pd)
}
