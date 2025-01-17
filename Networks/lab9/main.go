package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
)

type Server struct{
	ip             string
	port           string
	ipClient       string
	portClient     string
	from           string
	password       string
}

func NewServer(ip, port, ipClient, portClient string) *Server {
	return &Server{
		ip:          ip,
		port:        port,
		ipClient:    ipClient,
		portClient:  portClient,
	}
}

func (s *Server) httpHandlerForSendEmails(w http.ResponseWriter, r *http.Request) {
	from := s.from
	password := s.password
	to := r.FormValue("to")
	subject := r.FormValue("subject")
	message := r.FormValue("message")

	fmt.Println(subject)
	fmt.Println(message)

	s.SendEmail(from, password, to, subject, message)
}

func (s *Server) httpHandlerForAuth(w http.ResponseWriter, r *http.Request) {
	from := r.FormValue("from")
	password := r.FormValue("password")

	s.from = from
	s.password = password

	http.Redirect(w, r, "/client", http.StatusSeeOther)

	fmt.Println(from)
	fmt.Println(password)
}

func (s *Server) SendEmail(from string, password string, toList string, subject string, message string) {
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%s", s.ipClient, s.portClient)) // устанавливаем соединение
	if err != nil {
		log.Printf("Error connecting to next peer: %v\n", err)
		return
	}

	defer conn.Close() // закрываем соединение в конце работы функции

	newMessage := fmt.Sprintf("%s %s %s %s %s", from, toList, password, subject, message)

	fmt.Println(newMessage)

	_, err = fmt.Fprintf(conn, "%s\n", newMessage) // отправляем message слудующему пиру
	if err != nil {
		log.Printf("Error sending message: %v\n", err)
	}
}

func main() {
	if len(os.Args) != 5 {
		fmt.Printf("Usage: %s <ip> <port>\n", os.Args[0])
		return
	}

	ip := os.Args[1]
	port := os.Args[2]
	ipClient := os.Args[3]
	portClient := os.Args[4]

	server := NewServer(ip, port, ipClient, portClient)


	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        http.ServeFile(w, r, "indexstart.html")
    })

	http.HandleFunc("/client", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})

	http.HandleFunc("/auth", server.httpHandlerForAuth)

	http.HandleFunc("/send", server.httpHandlerForSendEmails)

	addr := fmt.Sprintf("%s:%s", server.ip, server.port)

	fmt.Printf("Start server on http://%s\n", addr)
	http.ListenAndServe(addr, nil)
}
