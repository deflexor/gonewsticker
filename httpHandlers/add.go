package httpHandlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"log"

	"github.com/deflexor/gonewsticker/httpHandlers/httpUtils"
	"github.com/deflexor/gonewsticker/structs"
)

func Add(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	byteData, err := ioutil.ReadAll(r.Body)

	if err != nil {
		httpUtils.HandleError(&w, 500, "Internal Server Error", "Error reading data from body", err)
		return
	}

	var message structs.NewsMessage

	err = json.Unmarshal(byteData, &message)

	if err != nil {
		httpUtils.HandleError(&w, 500, "Internal Server Error", "Error unmarhsalling JSON", err)
		return
	}

	if message.Text == "" || message.Sender == "" {
		httpUtils.HandleError(&w, 400, "Bad Request", "Unmarshalled JSON didn't have required fields", nil)
		return
	}

	// id := storage.Add(message)
	id := 1

	log.Println("Added message:", message)

	httpUtils.HandleSuccess(&w, structs.ID{ID: id})
}
