package main

import (
	"time"
	"github.com/deflexor/gonewsticker/storage"
	"sort"
	"github.com/deflexor/gonewsticker/httpHandlers"
	"github.com/deflexor/gonewsticker/structs"
	"log"
	"net/http"
	"strconv"
	"github.com/ungerik/go-rss"
)

const PORT = 8080
var NEWS_URLS = []string{ "https://hi-tech.mail.ru/rss/all/", "https://news.yandex.ru/science.rss" }
var NEWS_DATE_FMTS = []string { time.RFC1123Z, "02 Jan 2006 15:04:05 -0700" }

// var messageId = 0

type RSSFetcher interface {
	Get(url string) (resp *rss.Channel, err error)
}

type rssFetchGoRSS struct {}

func (f rssFetchGoRSS) Get(url string) (resp *rss.Channel, err error) {
	return rss.Read(url)
}

func FetchNewsSlice (url string, fmt string, resultC chan []structs.NewsMessage, fetcher RSSFetcher) {
	var news []structs.NewsMessage
	rssChannel, err := fetcher.Get(url)
	if err != nil {
		log.Println(err)
		resultC <- news
		return
	}
	// rssChannel.Item = rssChannel.Item[:2]
	for _, item := range rssChannel.Item {
	  created, err := item.PubDate.ParseWithFormat(fmt)
	  if err != nil {
		  continue
	  }
	  news = append(news, structs.NewsMessage{
		  	Sender: url,
			GUID: item.GUID,
			URL: item.Link,
			Title: item.Title,
			Text:  item.FullText,
			Created: created })
	}
	resultC <- news
}

func FetchNews(fetcher RSSFetcher) {
	log.Print("Fetching news..")
	var news []structs.NewsMessage
	resultC := make(chan []structs.NewsMessage)
	for i, url := range NEWS_URLS {
		go FetchNewsSlice(url, NEWS_DATE_FMTS[i], resultC, fetcher)
	}
	for range NEWS_URLS {
		n := <-resultC
		news = append(news, n...)
	}
	sort.Slice(news, func (i, j int) bool { return news[i].Created.After(news[j].Created) })
	log.Printf("got news: %d", len(news))
	storage.AddMany(news)
	log.Println("Done.")
}

func NewsFetcher() {
	for {
		FetchNews(rssFetchGoRSS{})
		time.Sleep(time.Hour * 2)
	}
}

func main() {
	go NewsFetcher()

	log.Println("started HTTP Server.")

	http.HandleFunc("/", httpHandlers.HandleRequest)

	var err = http.ListenAndServe(":"+strconv.Itoa(PORT), nil)

	if err != nil {
		log.Panicf("Server failed starting. Error: %s\n", err)
	}
}
