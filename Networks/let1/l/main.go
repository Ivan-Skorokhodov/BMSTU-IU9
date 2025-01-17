package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"strings"
)

func main() {
	value := "46304715bc92ea35239ced013f319d2f"
	url := "http://pstgu.yss.su/iu9/networks/let1_2024/getkey.php?hash=" + value

	// Отправляем GET-запрос
	response, err := http.Get(url)
	if err != nil {
		log.Fatalf("Ошибка при отправке запроса: %v", err)
	}
	defer response.Body.Close()

	scanner := bufio.NewScanner(response.Body)
	
	for scanner.Scan() {
		l := scanner.Text()
		parts := strings.Split(l, ": ")
		password := parts[1]
		fmt.Println(password)

		url := "http://pstgu.yss.su/iu9/networks/let1_2024/send_from_go.php?subject=let1 ИУ9-31Б Скороходов Иван&fio=Скороходов Иван&pass=" + password
		response, err := http.Get(url)
		if err != nil {
			log.Fatalf("Ошибка при отправке запроса: %v", err)
		}
		defer response.Body.Close()
	}
}