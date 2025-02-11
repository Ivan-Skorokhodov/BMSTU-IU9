package main

import (
	"fmt"
	"log"

	"github.com/gorilla/websocket"
)

func main() {
	serverURL := "ws://localhost:8080/ws"
	conn, _, err := websocket.DefaultDialer.Dial(serverURL, nil)
	if err != nil {
		log.Fatal("Ошибка подключения к серверу:", err)
	}
	defer conn.Close()

	fmt.Println("Подключено к серверу WebSocket.")

	input := 10
	response := fmt.Sprintf("%d", input)
	err = conn.WriteMessage(websocket.TextMessage, []byte(response))
	if err != nil {
		log.Println("Ошибка отправки сообщения серверу:", err)
	}

	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("Ошибка получения сообщения от сервера:", err)
			break
		}
		
		fmt.Printf("Ответ от сервера: %s\n", message)

		err = conn.WriteMessage(websocket.TextMessage, []byte(message))
		if err != nil {
			log.Println("Ошибка отправки сообщения серверу:", err)
		}
	}

}
