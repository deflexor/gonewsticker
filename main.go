package main

import (
	"github.com/deflexor/gonewsticker/storage"
	"sort"
//	"io/ioutil"
	"github.com/deflexor/gonewsticker/httpHandlers"
//	"github.com/deflexor/gonewsticker/httpHandlers/httpUtils"
	"github.com/deflexor/gonewsticker/structs"
	"log"
	"net/http"
	"strconv"
	"github.com/ungerik/go-rss"
)

const PORT = 8080
var NEWS_URLS = []string{ "https://hi-tech.mail.ru/rss/all/", "https://news.yandex.ru/science.rss" }

var messageId = 0

func fetchNews() {
	var news []structs.NewsMessage
	for _, url := range NEWS_URLS {
		channel, err := rss.Read(url)
		if err != nil {
			log.Println(err)
		}
	
		log.Println(channel.Title)
	
		for _, item := range channel.Item {
		  created, err := item.PubDate.Parse()
		  if err != nil {
			  log.Printf("%v", err)
			  continue
		  }
		  news = append(news, structs.NewsMessage{
				GUID: item.GUID,
				Title: item.Title,
				Text:  item.FullText,
				Created: created })
		}
		// log.Printf("%s", rss)
	}
	sort.Slice(news, func (i, j int) bool { return news[i].Created.After(news[j].Created) })
	storage.AddMany(news)
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
