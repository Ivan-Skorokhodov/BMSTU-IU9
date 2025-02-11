package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func websocketHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Ошибка обновления до WebSocket:", err)
		return
	}
	defer conn.Close()

	log.Println("WebSocket-соединение установлено.")

	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("Ошибка чтения сообщения:", err)
			break
		}

		log.Printf("Сообщение от клиента: %s\n", message)

		num, err := strconv.Atoi(string(message))
		if err != nil {
			log.Println("Ошибка преобразования в число:", err)
			continue
		}

		log.Printf("Получено число: %d\n", num)

		num++

		response := fmt.Sprintf("%d", num)
		err = conn.WriteMessage(websocket.TextMessage, []byte(response))
		if err != nil {
			log.Println("Ошибка отправки сообщения:", err)
			break
		}
	}
}

func main() {
	http.HandleFunc("/ws", websocketHandler)

	fmt.Println("Сервер запущен на http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
