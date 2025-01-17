package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"net/http"
	"net/smtp"
	"os"
	"strings"
	"time"

	"github.com/gorilla/websocket"
)

type Server struct{
	ip         string
	WSport     string
	TCPport    string
	wsConn     *websocket.Conn
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

func NewServer(ip, WSport, TCPport string) *Server {
	return &Server{
		ip:          ip,
		WSport:      WSport,
		TCPport:     TCPport,
	}
}

func (s *Server) SendEmail(from string, password string, toList []string, subject, message string){
	host := "smtp.mail.ru"
	port := "587"

	auth := smtp.PlainAuth("", from, password, host)

	for i := 0; i < len(toList); i++ {
		line := fmt.Sprintf("send mail to: " + toList[i])

		fmt.Println(line)
		if err := s.wsConn.WriteJSON(line); err != nil {
			log.Println("Ошибка отправки данных:", err)
			return
		}

		subject := subject

		msg := message
		/*
		"<p>Уважаемый " + toList[i] + ",</p>\n" +
        "<p>Это письмо отправлено Скороходовым Иваном.</p>\n" +
        "<p>C уважением,</p>\n" +
        "<p>Скороходов Иван</p>" +
        "<p>небольшой бонус - рецепт приготовления пельменей:</p>\n" +
        "<p>Ингредиенты:</p>\n" +
        "<ul>\n" +
        "<li>500 муки</li>\n" +
        "<li>250 воды</li>\n" +
        "<li>1 яйцо</li>\n" +
        "<li>соль, перец</li>\n" +
        "<li>начинка (мясо, лук, чеснок)</li>\n" +
        "</ul>\n" +
        "<p>Инструкция:</p>\n" +
        "<ol>\n" +
        "<li>Замесить тесто из муки, воды, яйца и соли.</li>\n" +
        "<li>Раскатать тесто тонким слоем.</li>\n" +
        "<li>Сделать начинку из мяса, лука и чеснока.</li>\n" +
        "<li>Сложить пельмени и варить в кипящей воде 10-15 минут.</li>\n" +
        "</ol>\n" +
        "<p>Приятного аппетита!</p>"
		*/

		msgRes := fmt.Sprintf("To: %s\r\nFrom: %s\r\nSubject: %s\r\n", toList[i], from, subject)
		msgRes += "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\r\n"
		msgRes += fmt.Sprintf("\r\n%s\r\n", msg)

		err := smtp.SendMail(host+":"+port, auth, from, []string{toList[i]}, []byte(msgRes))

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		time.Sleep(2*time.Second)
	}

	line := "Successfully sent mail to all users in toList"

	fmt.Println(line)

	if err := s.wsConn.WriteJSON(line); err != nil {
		log.Println("Ошибка отправки данных:", err)
		return
	}
}

func (s *Server) handleTCPConnection(TCPconn net.Conn) {
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

		from := args[0]
		toList := args[1]
		password := args[2]
		subject := args[3]
		message = strings.Join(args[4:], " ")
		s.SendEmail(from, password, []string{toList}, subject, message)
	}
}

func (s *Server) startTCPServer() {
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%s", s.ip, s.TCPport))
	if err != nil {
		log.Fatalf("Error starting server: %v\n", err)
	}
	defer listener.Close()

	fmt.Printf("Email-Client listening on %s:%s\n", s.ip, s.TCPport)

	for {
		TCPConn, err := listener.Accept()
		if err != nil {
			log.Printf("Error accepting connection: %v\n", err)
			continue
		}
		go s.handleTCPConnection(TCPConn)
	}
}

func (s *Server) handleWebSocket(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Ошибка обновления до WebSocket:", err)
		return
	}
	defer conn.Close()

	s.wsConn = conn
	go s.startTCPServer() 

	if err := conn.WriteJSON("Hello, let's send emails for someone! (Skorokhodov's software)"); err != nil {
		log.Println("Ошибка отправки данных:", err)
		return
	}

	reader := bufio.NewReader(os.Stdin)
	for {
		command, _ := reader.ReadString('\n')
		command = strings.TrimSpace(command)

		switch {
		case command == "exit":
			os.Exit(0)
		default:
			fmt.Println("Unknown command. Available commands: print list, hide <peer_nunber>, show <peer_number>")
		}
	}
}


func main() {
	if len(os.Args) != 4 {
		fmt.Printf("Usage: %s <ip> <WSport> <TCPport>\n", os.Args[0])
		return
	}

	ip := os.Args[1]
	WSport := os.Args[2]
	TCPport := os.Args[3]

	server := NewServer(ip, WSport, TCPport)

	http.HandleFunc("/ws", server.handleWebSocket)

	addr := fmt.Sprintf("%s:%s", server.ip, server.WSport)
	fmt.Printf("Start server on http://%s\n", addr)
	http.ListenAndServe(addr, nil)
}
