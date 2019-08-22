package main

import (
	"io/ioutil"
	"github.com/deflexor/gonewsticker/httpHandlers"
//	"github.com/deflexor/gonewsticker/httpHandlers/httpUtils"
//	"github.com/deflexor/gonewsticker/structs"
	"log"
	"net/http"
	"strconv"
)

const PORT = 8080
var NEWS_URLS = []string{ "https://hi-tech.mail.ru/rss/all/", "https://news.yandex.ru/science.rss" }

var messageId = 0

func fetchNews() {
	
	for _, url := range NEWS_URLS {
		res, err := http.Get(url)
		if err != nil {
			log.Fatal(err)
		}
		rss, err := ioutil.ReadAll(res.Body)
		res.Body.Close()
		if err != nil {
			log.Fatal(err)
		}
		// fmt.Printf("%s", rss)
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
