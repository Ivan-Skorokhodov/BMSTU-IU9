package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/SlyMarbo/rss"
	_ "github.com/go-sql-driver/mysql"
)

const dsn = "iu9networkslabs:Je2dTYr6@tcp(students.yss.su:3306)/iu9networkslabs"

func main() {
    db, err := sql.Open("mysql", dsn)
    if err != nil {
        log.Fatal("Ошибка подключения к базе данных:", err)
    }
    defer db.Close()

	updateNews(db)
}

func updateNews(db *sql.DB) {
    feed, err := rss.Fetch("https://lenta.ru/rss")
	if err != nil {
		fmt.Println(err)
	}

    existingTitles, err := getExistingTitles(db)
    if err != nil {
        log.Println("Ошибка получения заголовков:", err)
        return
    }

    for _, item := range feed.Items {
        if _, exists := existingTitles[item.Title]; !exists {
            if _, err := db.Exec(`INSERT INTO iu9Skorokhodov (title, link, date) VALUES (?, ?, ?)`,
                item.Title, item.Link, time.Now()); err != nil {
                log.Println("Ошибка вставки новости:", err)
            } else {
                log.Printf("Добавлена новость: %s\n", item.Title)
            }
        }
    }
}

func getExistingTitles(db *sql.DB) (map[string]bool, error) {
    rows, err := db.Query(`SELECT title FROM iu9Skorokhodov`)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    titles := make(map[string]bool)
    for rows.Next() {
        var title string
        if err := rows.Scan(&title); err != nil {
            return nil, err
        }
        titles[title] = true
    }
    return titles, nil
}