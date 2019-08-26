package httpHandlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/deflexor/gonewsticker/storage"

	"log"

	"github.com/deflexor/gonewsticker/httpHandlers/httpUtils"
	"github.com/deflexor/gonewsticker/structs"
)

func AddComment(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	byteData, err := ioutil.ReadAll(r.Body)
	if err != nil {
		httpUtils.HandleError(&w, 500, "Internal Server Error", "", err)
		return
	}

	guid, ok := r.URL.Query()["guid"]
	if !ok && len(guid[0]) < 1 {
		httpUtils.HandleError(&w, 400, "Ошибка при добавлении комментария!", "", nil)
		return
	}

	var c structs.Comment

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
