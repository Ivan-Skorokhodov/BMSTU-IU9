package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"

	"github.com/gorilla/websocket"
)

var WSConn *websocket.Conn

//описание структуры пира
type Peer struct {
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
func NewPeer(peerNumber, ip, HTTPport, TCPport, WSPort, nextPeerIP, nextPeerPort string) *Peer {
	return &Peer{
		peerNumber:    peerNumber,
		ip:            ip,
		HTTPport:      HTTPport,
		TCPport:       TCPport,
		WSPort:        WSPort,
		nextPeerIP:    nextPeerIP,
		nextPeerPort:  nextPeerPort,
		visiblePeers:  make(map[string]bool), // все пиры изначально видны, их количество захардкожено, всего 3 пира (0, 1, 2)
	}
}

// переменная для WebSocket обработчика
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// Управление видимостью пиров
func (p *Peer) manageVisibility(command string, WSConn *websocket.Conn) {
	args := strings.Split(command, " ")
	if len(args) != 3 {
		fmt.Println("Invalid command. Use: hide <peer_number> or show <peer_number>")
		return
	}

	p.mu.Lock()
	defer p.mu.Unlock()

	peerNumber := args[1]
	count, _ := strconv.Atoi(args[2])
	if args[0] == "hide" {
		p.visiblePeers[peerNumber] = false
		fmt.Printf("%s is now hidden\n", peerNumber)

		count++
		strCount := strconv.Itoa(count)
		message := fmt.Sprintf("hide %s %s", peerNumber, strCount)
		p.sendMessageToPeer(message)

		p.sendMessageToClient(fmt.Sprintf("%s is now hidden\n", peerNumber), WSConn)

	} else if args[0] == "show" {
		p.visiblePeers[peerNumber] = true
		fmt.Printf("%s is now visible\n", peerNumber)

		count++
		strCount := strconv.Itoa(count)
		message := fmt.Sprintf("show %s %s", peerNumber, strCount)
		p.sendMessageToPeer(message)

		p.sendMessageToClient(fmt.Sprintf("%s is now visible\n", peerNumber), WSConn)

	} else {
		fmt.Println("Unknown command. Use: hide or show")
	}
}

// Функция обработки входящих сообщений с пиров
func (p *Peer) handleTCPConnection(TCPconn net.Conn, WSConn *websocket.Conn) {
	defer TCPconn.Close()
	reader := bufio.NewReader(TCPconn)
	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			return
		}
		message = strings.TrimSpace(message)
		fmt.Printf("From server interface: Received message: %s\n", message)

		args := strings.Split(message, " ")
		if len(args) != 3 {
			fmt.Println("Invalid command. Use: hide <peer_number> or show <peer_number>")
			return
		}

		count, _ := strconv.Atoi(args[2])

		if count < 3 {
			p.manageVisibility(message, WSConn)
		}
	}
}

// Запуск TCP сервера для обработки входящих соединений
func (p *Peer) startTCPServer(WSConn *websocket.Conn) {
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%s", p.ip, p.TCPport))
	if err != nil {
		log.Fatalf("Error starting server: %v\n", err)
	}
	defer listener.Close()

	fmt.Printf("Peer %s listening on %s:%s\n", p.peerNumber, p.ip, p.TCPport)

	for {
		TCPConn, err := listener.Accept()
		if err != nil {
			log.Printf("Error accepting connection: %v\n", err)
			continue
		}
		go p.handleTCPConnection(TCPConn, WSConn)
	}
}

// Вывод списка видимых пиров (пока что только в консоль)
func (p *Peer) printVisiblePeers(WSConn *websocket.Conn) {
	p.mu.Lock()
	defer p.mu.Unlock()

	message := ""

	for peer, visible := range p.visiblePeers {
		if visible {
			str := fmt.Sprintf("%s peer is visible\n", peer)
			message += str
			fmt.Print(str)
		}
	}
	
	p.sendMessageToClient(message, WSConn)
}

// Отправка сообщения следующему пиру
func (p *Peer) sendMessageToPeer(message string) {
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%s", p.nextPeerIP, p.nextPeerPort)) // устанавливаем соединение
	if err != nil {
		log.Printf("Error connecting to next peer: %v\n", err)
		return
	}

	defer conn.Close() // закрываем соединение в конце работы функции

	_, err = fmt.Fprintf(conn, "%s\n", message) // отправляем message слудующему пиру
	if err != nil {
		log.Printf("Error sending message: %v\n", err)
	}
}

func (p *Peer) sendMessageToClient(message string, WSConn *websocket.Conn) {
	err := WSConn.WriteMessage(websocket.TextMessage, []byte(message))
	if err != nil {
		log.Printf("Error sending message to client: %v\n", err)
		return
	}
}

// WebSocket обработчик
func (p *Peer) wsHandler(w http.ResponseWriter, r *http.Request) {
	WSConn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("Error upgrading to websocket: %v\n", err)
		return
	}
	defer WSConn.Close()

	go p.startTCPServer(WSConn)      // Запуск TCP сервера

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("write command> ")
		command, _ := reader.ReadString('\n')
		command = strings.TrimSpace(command)

		switch {
		case command == "print list":
			p.printVisiblePeers(WSConn)
		case strings.HasPrefix(command, "hide"), strings.HasPrefix(command, "show"):
			p.manageVisibility(command, WSConn)
		default:
			fmt.Println("Unknown command. Available commands: print list, hide <peer_nunber>, show <peer_number>")
		}
	}
}

// Запуск HTTP сервера для WebSocket
func (p *Peer) startWebSocketServer() {
	http.HandleFunc("/ws", p.wsHandler)
	fmt.Printf("WebSocket Server is listening on %s:%s\n", p.ip, p.WSPort)
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", p.ip, p.WSPort), nil))
}

// Запуск пира и обработка команд
func (p *Peer) start() {
	p.visiblePeers[strconv.Itoa(0)] = true
	p.visiblePeers[strconv.Itoa(1)] = true
	p.visiblePeers[strconv.Itoa(2)] = true
	
	p.startWebSocketServer() // Запуск WebSocket сервера
}

func main() {
	if len(os.Args) != 8 {
		fmt.Printf("Usage: %s <peerNumber> <ip> <HTTP_port> <TCP_port> <WS_port> <next_peer_ip> <next_peer_port>\n", os.Args[0])
		return
	}

	peerNumber := os.Args[1]
	ip := os.Args[2]
	HTTPport := os.Args[3]
	TCPport := os.Args[4]
	WSPort := os.Args[5]
	nextPeerIP := os.Args[6]
	nextPeerPort := os.Args[7]

	peer := NewPeer(peerNumber, ip, HTTPport, TCPport, WSPort, nextPeerIP, nextPeerPort)
	peer.start()
}