package main

import (
	"fmt"
	"html/template"
	"net/http"

	log "github.com/mgutz/logxi/v1"
	"golang.org/x/net/html"
)

const INDEX_HTML = `
    <!doctype html>
    <html lang="ru">
        <head>
            <meta charset="utf-8">
            <title>Last news from lenta.ru</title>
        </head>
        <body>
            {{if .}}
                {{range .}}
                    <a href="{{.Ref}}">{{.Time}}</a>
                    <img src={{.Title}}>
                    <br/>
                {{end}}
            {{else}}
                Error of downloading news!
            {{end}}
        </body>
    </html>
    `

var indexHtml = template.Must(template.New("index").Parse(INDEX_HTML))

func serveClient(response http.ResponseWriter, request *http.Request) {
	path := request.URL.Path
	fmt.Println("got request", "Method", request.Method, "Path", path)
	if path != "/" && path != "/index.html" {
		log.Error("invalid path", "Path", path)
		response.WriteHeader(http.StatusNotFound)
	} else if err := indexHtml.Execute(response, downloadNews()); err != nil {
		log.Error("HTML creation failed", "error", err)
	} else {
		fmt.Println("response sent to client successfully")
	}
}

func main() {
	http.HandleFunc("/", serveClient)
	fmt.Println("starting listener")
	log.Error("listener failed", "error", http.ListenAndServe("127.0.0.1:9090", nil))
}

func getAttr(node *html.Node, key string) string {
	for _, attr := range node.Attr {
		if attr.Key == key {
			return attr.Val
		}
	}
	return ""
}

func getChildren(node *html.Node) []*html.Node {
	var children []*html.Node
	for c := node.FirstChild; c != nil; c = c.NextSibling {
		children = append(children, c)
	}
	return children
}

func isElem(node *html.Node, tag string) bool {
	return node != nil && node.Type == html.ElementNode && node.Data == tag
}

func isClass(node *html.Node, tag string) bool{
	return node != nil && node.Type == html.ElementNode && getAttr(node, "class") == tag
}

func isText(node *html.Node) bool {
	return node != nil && node.Type == html.TextNode
}

func isDiv(node *html.Node, class string) bool {
	return isElem(node, "div") && getAttr(node, "class") == class
}

type Item struct {
	Ref, Time, Title string
}
func readItem(item *html.Node) *Item {
/*
	fmt.Println("------")
	fmt.Println(item.Data)
    fmt.Println()
	fmt.Println(item.FirstChild.Attr)
    fmt.Println()
	fmt.Println(item.FirstChild.FirstChild.FirstChild.Attr)
	fmt.Println("------")
*/
	a := item.FirstChild;
	h3 := a.FirstChild.FirstChild;
	title := h3.Data
	_ = title
	time := a.FirstChild.NextSibling.FirstChild.FirstChild.Data

	return &Item{
		Ref: getAttr(item.FirstChild, "href"), //ссылка
		Time: time, //заголовок
		Title: getAttr(item.FirstChild.FirstChild.FirstChild, "srcset"),
	}
}

func search(node *html.Node) []*Item {
	if isDiv(node, "zT5wwAPN fQtJ19Ei") {
		var items []*Item
		for c := node.FirstChild; c != nil; c = c.NextSibling {
			if isClass(c, "XSvLK2D0 abGoxuyb") {
				if item := readItem(c); item != nil {
					items = append(items, item)
				}
			}
		}
		return items
	}
	for c := node.FirstChild; c != nil; c = c.NextSibling {
		if items := search(c); items != nil {
			return items
		}
	}
	return nil
}

func downloadNews() []*Item {
	fmt.Println("sending request to news.rambler.ru")
	if response, err := http.Get("https://news.rambler.ru/latest/"); err != nil {
		log.Error("request to news.rambler.ru failed", "error", err)
	} else {
		defer response.Body.Close()
		status := response.StatusCode
		fmt.Println("got response from news.rambler.ru", "status", status)
		if status == http.StatusOK {
/*
			fmt.Println(response)
			fmt.Println("-------")
*/
			if doc, err := html.Parse(response.Body); err != nil {
				log.Error("invalid HTML from news.rambler.ru", "error", err)
			} else {
/*
				fmt.Println(doc)
				fmt.Println("-------")
				fmt.Println(response.Body)
*/
				fmt.Println("HTML from news.rambler.ru parsed successfully")
				return search(doc)
			}
		}
	}
	return nil
}