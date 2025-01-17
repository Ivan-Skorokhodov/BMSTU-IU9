package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"sync"
)

//описание структуры пира
type Server struct {
	peerNumber    string
	ip            string
	HTTPport     string
	TCPport       string
	WSPort        string
	nextPeerIP    string
	nextPeerPort  string
	visiblePeers  map[string]bool
	mu            sync.Mutex
}

// Создание нового пира
func NewServer(ip, HTTPport, nextPeerIP, nextPeerPort string) *Server {
	return &Server{
		ip:            ip,
		HTTPport:      HTTPport,
		nextPeerIP:    nextPeerIP,
		nextPeerPort:  nextPeerPort,
	}
}

func (s *Server) httpHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "~/Skorokhodov/httpServer/test.html")
}

func (s *Server) httpHandlerAdder(w http.ResponseWriter, r *http.Request) {
	operation := r.FormValue("operation")
	peerNumber := r.FormValue("PeerNumber")

	message := fmt.Sprintf("%s %s 0", operation, peerNumber)

	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%s", s.nextPeerIP, s.nextPeerPort)) // устанавливаем соединение
	if err != nil {
		log.Printf("Error connecting to next peer: %v\n", err)
		return
	}

	defer conn.Close() // закрываем соединение в конце работы функции

	_, err = fmt.Fprintf(conn, "%s\n", message) // отправляем message слудующему пиру
	if err != nil {
		log.Printf("Error sending message: %v\n", err)
	}

	fmt.Println(message)
}

func (s *Server) startHTTPServer() {
	http.HandleFunc("/", s.httpHandler)
	http.HandleFunc("/add", s.httpHandlerAdder)
	fmt.Printf("HTTP Server is listening on %s:%s\n", s.ip, s.HTTPport)
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", s.ip, s.HTTPport), nil))
}


func main() {
	if len(os.Args) != 5 {
		fmt.Printf("Usage: %s <ip> <HTTP_port> <next_peer_ip> <next_peer_port>\n", os.Args[0])
		return
	}

	ip := os.Args[1]
	HTTPport := os.Args[2]
	nextPeerIP := os.Args[3]
	nextPeerPort := os.Args[4]

	server := NewServer(ip, HTTPport, nextPeerIP, nextPeerPort)
	server.startHTTPServer()
}