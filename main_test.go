package main

import (
	"errors"
	"testing"
	"github.com/deflexor/gonewsticker/storage"
	"github.com/ungerik/go-rss"
)


var testData = []rss.Channel{
	rss.Channel{
		Title:"Яндекс.Новости: Наука", 
		Link:"https://news.yandex.ru/science.html?from=rss", 
		Description:"Первая в России служба автоматической обработки и системати...",
		Language:"",
		LastBuildDate:"23 Aug 2019 17:43:38 +0000",
		Item: []rss.Item{
			rss.Item{
				Title:"Теломеры помогли мышиным клеткам остаться стволовыми",
				Link:"https://news.yandex.ru/story/Telomery_pomogli_myshinym_kletkam_ostatsya_stvolovymi--2bf6b6f0002ccdc454a9c6477a8494bf?lang=ru&from=rss&stid=Q9Ec",
				Comments:"",
				PubDate:"23 Aug 2019 16:22:18 +0000",
				GUID:"https://news.yandex.ru/story/Telomery_pomogli_myshinym_kletkam_ostatsya_stvolovymi--2bf6b6f0002ccdc454a9c6477a8494bf?lang=ru&from=rss&stid=Q9Ec",
				Category: []string(nil),
				Enclosure:[]rss.ItemEnclosure(nil),
				Description:"Испанские исследователи обнаружили белок, который не только..",
				Author:"",
				Content:"",
				FullText:""},
			rss.Item{
				Title:"Ученые создали браслет, предсказывающий вспышки агрессии при аутизме",
				Link:"https://news.yandex.ru/story/Uchenye_sozdali_braslet_predskazyvayushhij",
				Comments:"",
				PubDate:"23 Aug 2019 06:05:00 +0000",
				GUID:"https://news.yandex.ru/story/Uchenye_sozdali_braslet_predskazyvayushhij",
				Category:[]string(nil),
				Enclosure:[]rss.ItemEnclosure(nil), 
				Description:"Американские ученые из Северо-Восточного университета 84 %.",
				Author:"",
				Content:"",
				FullText:""}}},
	rss.Channel{
		Title:"Яндекс.Новости: Наука", 
		Link:"https://news.yandex.ru/science.html?from=rss", 
		Description:"Первая в России служба автоматической обработки и системати...",
		Language:"",
		LastBuildDate:"23 Aug 2019 19:43:38 +0000",
		Item: []rss.Item{
			rss.Item{
				Title:"Теломеры помогли мышиным клеткам остаться стволовыми",
				Link:"https://news.yandex.ru/story/Telomery_pomogli_myshinym_kletkam_ostatsya_stvolovymi--2bf6b6f0002ccdc454a9c6477a8494bf?lang=ru&from=rss&stid=Q9Ec",
				Comments:"",
				PubDate:"23 Aug 2019 17:15:18 +0000",
				GUID:"https://news.yandex.ru/story/Telomery_pomogli_myshinym_kletkam_ostatsya_stvolovymi--2bf6b6f0002ccdc454a9c6477a8494bf?lang=ru&from=rss&stid=Q9Ec",
				Category: []string(nil),
				Enclosure:[]rss.ItemEnclosure(nil),
				Description:"Испанские исследователи обнаружили белок, который не только..",
				Author:"",
				Content:"",
				FullText:""},
			rss.Item{
				Title:"Ученые создали браслет, предсказывающий вспышки агрессии при аутизме",
				Link:"https://news.yandex.ru/story/Fresh",
				Comments:"",
				PubDate:"23 Aug 2019 06:05:00 +0000",
				GUID:"https://news.yandex.ru/story/Fresh",
				Category:[]string(nil),
				Enclosure:[]rss.ItemEnclosure(nil), 
				Description:"Самаая свежая новость.",
				Author:"",
				Content:"",
				FullText:""}}},
	rss.Channel{
		Title:"Hi-Tech Mail.ru / Публикации",
		Link:"https://hi-tech.mail.ru",
		Description:"Новости и полезная информация о цифровых...",
		Language:"ru",
		LastBuildDate:"Fri, 23 Aug 2019 21:04:57 +0300",
		Item:[]rss.Item{
			rss.Item{
				Title:"Представили зубную щетку для&nbsp;самых ленивых",
				Link:"https://hi-tech.mail.ru/news/encompass_zubnaya_shetka/",
				Comments:"",
				PubDate:"Fri, 23 Aug 2019 18:04:00 +0300",
				GUID:"ebb33fa2a8c2da2fab66edb922c2fae3",
				Category:[]string(nil),
				Enclosure:[]rss.ItemEnclosure{
					rss.ItemEnclosure{
						URL:"https://htstatic.imgsmail.ru/pic_share/2ecf662edb643f957dfd448330d68262/1668379/c/45275?time=1566581682",
						Type:"image/png"}},
				Description:"<p>Создатели новой автоматической зубной щетки обещают, что&nbsp;она сделает все за&nbsp;20 секунд.</p>",
				Author:"",
				Content:"",
				FullText:""},
			rss.Item{
				Title:"iPhone дорожают, потому что&nbsp;их&nbsp;не&nbsp;хотят покупать",
				Link:"https://hi-tech.mail.ru/news/nikto_ne_pokupaet_iphone_dorozhayut/",
				Comments:"",
				PubDate:"Fri, 23 Aug 2019 18:01:00 +0300",
				GUID:"028018e0d2c272445e3c1a7562a9eecd",
				Category:[]string(nil),
				Enclosure:[]rss.ItemEnclosure{},
				Description:"<p>Американские пользователи стали",
				Author:"",
				Content:"",
				FullText:""}}}}


type rssFetchTest struct {
	run int
}

func (f rssFetchTest) Get(url string) (resp *rss.Channel, err error) {
	if f.run == 2 {
		if url == "https://hi-tech.mail.ru/rss/all/" {
			return &testData[2], nil
		} else if url == "https://news.yandex.ru/science.rss" {
			return &testData[1], nil
		} else {
			return nil, errors.New("ouch")
		}
	}
	if url == "https://hi-tech.mail.ru/rss/all/" {
		return &testData[2], nil
	} else if url == "https://news.yandex.ru/science.rss" {
		return &testData[0], nil
	} else {
		return nil, errors.New("ouch")
	}
}

func TestFetchNews(t *testing.T) {
	storage.Clear()
	// fetch 1
	FetchNews(rssFetchTest{ run: 1 })
	news := storage.Get()
	if len(news) != 4 {
		t.Errorf("length of storage != 4 (%d)\n", len(news))
	}
	// fetch 2 (same data)
	FetchNews(rssFetchTest{ run: 1 })
	news = storage.Get()
	if len(news) != 4 {
		t.Errorf("length of storage != 4 (%d)\n", len(news))
	}
	// fetch 3 (1 new)
	FetchNews(rssFetchTest{ run: 2 })
	news = storage.Get()
	if len(news) != 5 {
		t.Errorf("length of storage != 5 (%d)\n", len(news))
	}
	/*for _, n := range news {
		t.Errorf("%s, %s", n.Sender, n.Created)
	}*/
}
