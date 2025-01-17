package main

// tcp dump -A -c 300 | grep key > out
// cat out

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strings"
)

func main() {
    // Создаем команду tcpdump
    tcpdumpCmd := exec.Command("tcpdump", "-A", "-c", "9000")

    // Создаем команду grep для фильтрации по ключевому слову "key"
    grepCmd := exec.Command("grep", "key")

    // Получаем stdout команды tcpdump и подключаем его к stdin команды grep
    pipeReader, pipeWriter := io.Pipe()
    tcpdumpCmd.Stdout = pipeWriter
    grepCmd.Stdin = pipeReader

    // Открываем файл out.txt для записи вывода grep
    file, err := os.Create("out.txt")
    if err != nil {
        log.Fatalf("Не удалось создать файл: %v", err)
    }
    defer file.Close()

    // Вывод команды grep направляем в файл
    grepCmd.Stdout = file

    // Запускаем команду tcpdump
    if err := tcpdumpCmd.Start(); err != nil {
        log.Fatalf("Ошибка при запуске tcpdump: %v", err)
    }

    // Запускаем команду grep
    if err := grepCmd.Start(); err != nil {
        log.Fatalf("Ошибка при запуске grep: %v", err)
    }

    // Закрываем writer после завершения работы tcpdump
    go func() {
        defer pipeWriter.Close()
        tcpdumpCmd.Wait()
    }()

    // Ожидаем завершения grep
    grepCmd.Wait()

    log.Println("Захват и фильтрация пакетов завершены, результат записан в файл out.txt.")

    // Открываем файл для поиска строк с именем Skorokhodov
    searchForKeyInFile("out.txt", "Skorohodov")
}

// Функция для поиска строк в файле, содержащих имя "Skorokhodov"
func searchForKeyInFile(filename, name string) {
    // Открываем файл для чтения
    file, err := os.Open(filename)
    if err != nil {
        log.Fatalf("Не удалось открыть файл: %v", err)
    }
    defer file.Close()

    // Чтение файла построчно
    reader := bufio.NewReader(file)
    for {
        line, err := reader.ReadString('\n')
        if err == io.EOF {
            break
        } else if err != nil {
            log.Fatalf("Ошибка чтения файла: %v", err)
        }

        // Проверяем, содержит ли строка искомое имя
        if strings.Contains(line, name) {
            fmt.Printf("Найдена строка для %s: %s", name, line)
			
			parts := strings.SplitN(line, ":", 2)
	
			if len(parts) < 2 {
				fmt.Println("Ключ не найден")
				return
			}
			
			// Обрезаем лишние пробелы и берем первое слово
			value := strings.Fields(strings.TrimSpace(parts[1]))[0]
			fmt.Printf("Значение ключа: %s\n", value)

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
				resp, err := http.Get(url)
				if err != nil {
					log.Fatalf("Ошибка при отправке запроса: %v", err)
				}
				defer resp.Body.Close()
			}
		}


    }
}