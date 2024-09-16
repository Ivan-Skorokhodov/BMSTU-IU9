package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"

	"github.com/SlyMarbo/rss"
)

type data struct {
	Title string
	Link  string
}

type data_list struct {
	List []data
}

func HomeRouterHandler(w http.ResponseWriter, r *http.Request) {
	feed, err := rss.Fetch("https://www.press-line.ru/feed")
	if err != nil {
		fmt.Println(err)
	}

	data_list := data_list{List: []data{}}

	for _, item := range feed.Items {

		data := data{
			Title: item.Title,
			Link:  item.Link,
		}

		data_list.List = append(data_list.List, data)
	}

	tmpl, _ := template.ParseFiles("start.html")
    tmpl.Execute(w, data_list)
}

func main() {
	http.HandleFunc("/", HomeRouterHandler) // установим роутер

	err := http.ListenAndServe(":9000", nil) // задаем слушать порт
	fmt.Println("Server started")
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
