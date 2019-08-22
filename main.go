package main

import (
	"github.com/deflexor/gonewsticker/httpHandlers"
	"github.com/deflexor/gonewsticker/structs"
	"log"
	"net/http"
	"strconv"
)

const PORT = 8080

var messageId = 0

func createMessage(message string, sender string) structs.NewsMessage {
	messageId++
	return structs.NewsMessage{
		ID:     messageId,
		Sender: sender,
		Text:   message,
	}
}

func main() {
	log.Println("Creating dummy messages")

	/*storage.Add(createMessage("Testing", "1234"))
	storage.Add(createMessage("Testing Again", "5678"))
	storage.Add(createMessage("Testing A Third Time", "9012"))
	*/
	log.Println("Attempting to start HTTP Server.")

	http.HandleFunc("/", httpHandlers.HandleRequest)

	var err = http.ListenAndServe(":"+strconv.Itoa(PORT), nil)

	if err != nil {
		log.Panicln("Server failed starting. Error: %s", err)
	}
}
