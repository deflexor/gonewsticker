package httpHandlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
	"log"

	"github.com/deflexor/gonewsticker/storage"
	"github.com/deflexor/gonewsticker/httpHandlers/httpUtils"
	"github.com/deflexor/gonewsticker/structs"
	"github.com/deflexor/gonewsticker/session"
)

func AddComment(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	byteData, err := ioutil.ReadAll(r.Body)
	if err != nil {
		httpUtils.HandleError(&w, 500, "Internal Server Error", "", err)
		return
	}
	ses, err := r.Cookie("SESSIONV2")
	if err != nil {
		httpUtils.HandleError(&w, 401, "Not authorized", "", err)
		return
	}
	id := ses.Value
	userData, err1 := session.GetUser(id)
	if err1 != nil {
		httpUtils.HandleError(&w, 401, "Not authorized", "", err)
		return
	}

	guid, ok := r.URL.Query()["guid"]
	if !ok && len(guid[0]) < 1 {
		httpUtils.HandleError(&w, 400, "Ошибка при добавлении комментария!", "", nil)
		return
	}

	var c structs.Comment

	c.Author = userData.PersonaName
	c.Avatar = userData.AvatarMedium

	err = json.Unmarshal(byteData, &c)

	if err != nil {
		httpUtils.HandleError(&w, 500, "Internal Server Error", "Error unmarhsalling JSON", err)
		return
	}

	if c.Text == "" || c.Author == "" {
		httpUtils.HandleError(&w, 400, "Не заполнены нужные поля!", "", nil)
		return
	}

	c.Added = time.Now()
	ok = storage.AddComment(guid[0], c)

	log.Println("Added comment:", c)

	httpUtils.HandleSuccess(&w, c)
}
