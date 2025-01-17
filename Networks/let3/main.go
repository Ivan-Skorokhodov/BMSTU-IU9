package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/knadh/go-pop3"
)

var countGlobal int

type Mail struct {
	Title string
}

type PageData struct {
	Mails []Mail
}

func mail() []Mail {
	p := pop3.New(pop3.Opt{
		Host:       "pop.mail.ru",
		Port:       995,
		TLSEnabled: true,
	})

	c, err := p.NewConn()
	if err != nil {
		log.Fatal(err)
	}
	defer c.Quit()

	if err := c.Auth("smtplabaskor@mail.ru", "6tutnurgQR3XyWbrhpha"); err != nil {
		log.Fatal(err)
	}

	count, _, _ := c.Stat()
	fmt.Println("total messages=", count)
	countGlobal = count
	var mails []Mail

	for id := 1; id <= count; id++ {
		m, _ := c.Retr(id)
		mails = append(mails, Mail{Title: string(m.Header.Get("subject"))})
		fmt.Println(id, "=", m.Header.Get("subject"))
	}

	return mails
}

func PrintHandler(w http.ResponseWriter, r *http.Request) {
	mails := mail()

	data := PageData{Mails: mails}
	tmpl, err := template.ParseFiles("start.html")
	if err != nil {
		log.Fatal(err)
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		log.Fatal(err)
	}
}

func DeleteHandler(w http.ResponseWriter, r *http.Request) {
	p := pop3.New(pop3.Opt{
		Host:       "pop.mail.ru",
		Port:       995,
		TLSEnabled: true,
	})

	c, err := p.NewConn()
	if err != nil {
		log.Fatal(err)
	}
	defer c.Quit()

	if err := c.Auth("smtplabaskor@mail.ru", "6tutnurgQR3XyWbrhpha"); err != nil {
		log.Fatal(err)
	}

	for id := 1; id <= countGlobal; id++ {
		c.Dele(id)
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}




func main() {

	mails := mail()
	fmt.Println(mails)

	http.HandleFunc("/", PrintHandler)
	http.HandleFunc("/delete", DeleteHandler)

	err := http.ListenAndServe(":9001", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

