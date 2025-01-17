package main

import (
	"net/http"

	log "github.com/mgutz/logxi/v1"
	"golang.org/x/net/html"
)

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
	fmt.Println(item.FirstChild.Attr)
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

/*
func readItem(item *html.Node) *Item {
	if a := item.FirstChild; isElem(a, "a") {
		if cs := getChildren(a); len(cs) == 2 && isElem(cs[0], "time") && isText(cs[1]) {
			return &Item{
				Ref:   getAttr(a, "href"),
				Time:  getAttr(cs[0], "title"),
				Title: cs[1].Data,
			}
		}
	}
	return nil
}

func search(node *html.Node) []*Item {
	if isDiv(node, "topnews__column") {

		fmt.Println("test isDiv topnews__column")

		var items []*Item
		for c := node.FirstChild; c != nil; c = c.NextSibling {

			if isDiv(c, "card-mini _topnews") {

				fmt.Println("test isDiv card-mini _topnews")

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
*/
func downloadNews() []*Item {
	log.Info("sending request to news.rambler.ru")
	if response, err := http.Get("https://news.rambler.ru/latest/"); err != nil {
		log.Error("request to news.rambler.ru failed", "error", err)
	} else {
		defer response.Body.Close()
		status := response.StatusCode
		log.Info("got response from news.rambler.ru", "status", status)
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
				log.Info("HTML from news.rambler.ru parsed successfully")
				return search(doc)
			}
		}
	}
	return nil
}
