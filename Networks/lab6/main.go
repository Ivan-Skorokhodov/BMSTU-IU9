package main

import (
	"database/sql"
	"log"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/websocket"
)

const dsn = "iu9networkslabs:Je2dTYr6@tcp(students.yss.su:3306)/iu9networkslabs"

type News struct {
    Title string
    Link  string
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

func main() {
    db, err := sql.Open("mysql", dsn)
    if err != nil {
        log.Fatal("Ошибка подключения к базе данных:", err)
    }
    defer db.Close()

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        http.ServeFile(w, r, "index.html")
    })

    http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		handleWebSocket(w, r, db)
	})

    log.Println("Сервер запущен на http://185.102.139.168:9090")
    http.ListenAndServe("185.102.139.168:9090", nil)
}

func handleWebSocket(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Ошибка обновления до WebSocket:", err)
		return
	}
	defer conn.Close()

	for {
		rows, err := db.Query(`SELECT title, link, date FROM iu9Skorokhodov`)
		if err != nil {
			log.Println("Ошибка запроса данных:", err)
			return
		}

		var newsItems []map[string]interface{}
		for rows.Next() {
			var title, link string
			var date []uint8
			if err := rows.Scan(&title, &link, &date); err != nil {
				log.Println("Ошибка сканирования строки:", err)
				return
			}
			newsItems = append(newsItems, map[string]interface{}{
				"title": title,
				"link":  link,
			})
		}
		rows.Close()

		if err := conn.WriteJSON(newsItems); err != nil {
			log.Println("Ошибка отправки данных:", err)
			return
		}

		time.Sleep(15 * time.Second)
	}
}