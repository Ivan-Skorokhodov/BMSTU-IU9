package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

type Server struct {
	ip              string
	HTTPport        string
	countCSSfiles	int
	countImagefiles int
}

func(s *Server) getImageFromWWW(url string, imageFile *os.File){
	response, err := http.Get(url)
    if err != nil {
        log.Fatalf("Ошибка при отправке запроса: %v", err)
    }
    defer response.Body.Close()

	fmt.Println(url)

	// Копируем содержимое ответа в файл
    _, err = io.Copy(imageFile, response.Body)
    if err != nil {
        log.Fatalf("Ошибка при создании файла: %v", err)
    }
}

// получаем CSS стили через url
func getCSSAndWriteToFile(url string, fileCSS *os.File){
	//делаем запрос к вебсайту
	response, err := http.Get(url)
	if err != nil {
		log.Fatalf("Ошибка при отправке запроса: %v", err)
	}

	//построчно записываем информацию о CSS стилях в fileCSS
	scanner := bufio.NewScanner(response.Body)
	defer response.Body.Close()

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Fprintln(fileCSS, line)
	}
}

//записываем измененные ссылкы на CSS в html строку
func changeUrlInCSSTag(line *string, webSite string, fileCSS *os.File){
	parts := strings.Split(*line, " ")
	for i, part := range parts { 
		// ссылки могут начинаться c http://, это надо обработать
		if strings.Contains(part, "href") && !(strings.Contains(part, "https") || strings.Contains(part, "http")){

			firstQuote := strings.Index(part, `"`)

			if firstQuote != -1 {
				secondQuote := strings.Index(part[firstQuote+1:], `"`) + firstQuote

				if secondQuote != 1 {
					urlCSS := webSite + "/" + part[firstQuote+1:secondQuote+1]
					getCSSAndWriteToFile(urlCSS, fileCSS)
					
					parts[i] = fmt.Sprintf("href=\"/%s\"", fileCSS.Name())
				}
			}

		}
	}

	*line = strings.Join(parts, " ")
}

//записываем измененные ссылки в html строку
func (s *Server) changeUrlInATag(line *string, webSite string) {
	parts := strings.Split(*line, " ")

	for i, part := range parts { 

		if strings.Contains(part, "href") && !(strings.Contains(part, "https") || strings.Contains(part, "http")){

			firstQuote := strings.Index(part, `"`)

			if firstQuote != -1 {
				secondQuote := strings.Index(part[firstQuote+1:], `"`) + firstQuote

				if secondQuote != 1 {
					parts[i] = fmt.Sprintf("href=\"http://%s:%s/?webSite=%s/%s%s", s.ip, s.HTTPport, webSite, part[firstQuote+1:secondQuote+1], part[secondQuote+1:])
				}
			}

		} else if strings.Contains(part, "href") {
			firstQuote := strings.Index(part, `"`)

			if firstQuote != -1 {
				secondQuote := strings.Index(part[firstQuote+1:], `"`) + firstQuote

				if secondQuote != 1 {
					parts[i] = fmt.Sprintf("href=\"http://%s:%s/?webSite=%s%s", s.ip, s.HTTPport, part[firstQuote+1:secondQuote+1], part[secondQuote+1:])
				}
			}

		}
	}

	*line = strings.Join(parts, " ")
}

func (s *Server) getUrlFromWWW(url, webSite string){
	//url := "http://www.gnuplot.info/"

	//делаем запрос к вебсайту
	response, err := http.Get(url)
	if err != nil {
		log.Fatalf("Ошибка при отправке запроса: %v", err)
	}

	//создаем файл out.html, который отдадим браузеру клиента
	file, err := os.Create("out.html")
    if err != nil {
        log.Fatalf("Ошибка при создании файла: %v", err)
    }
    defer file.Close()

	//начинаем парсить полученный с вебсайта htnl документ 
	scanner := bufio.NewScanner(response.Body)
	defer response.Body.Close()

	for scanner.Scan() {
		line := scanner.Text()

		//обрабатываем для CSS
		if strings.Contains(line, "link") && strings.Contains(line, ".css"){

			//создаем файл для CSS
			fileCSS, err := os.Create(fmt.Sprintf("./static/CSS%d.css", s.countCSSfiles))
			s.countCSSfiles++
			if err != nil {
				log.Fatalf("Ошибка при создании файла: %v", err)
			}
			defer fileCSS.Close()

			changeUrlInCSSTag(&line, webSite, fileCSS)
			fmt.Fprintln(file, line)

			continue
		}

		if strings.Contains(line, "href") && strings.Contains(line, "<a"){
			s.changeUrlInATag(&line, webSite)
			fmt.Fprintln(file, line)

			continue
		}

		if strings.Contains(line, "img"){

			imageFile, err := os.Create(fmt.Sprintf("./static/Img%d.png", s.countImagefiles))
			if err != nil {
				log.Fatalf("Ошибка при создании файла: %v", err)
			}
			s.countImagefiles++
			defer imageFile.Close()

			parts := strings.Split(line, " ")

			for i, part := range parts { 

				if strings.Contains(part, "src") && !(strings.Contains(part, "https") || strings.Contains(part, "http")){
		
					firstQuote := strings.Index(part, `"`)
		
					if firstQuote != -1 {
						secondQuote := strings.Index(part[firstQuote+1:], `"`) + firstQuote
		
						if secondQuote != 1 {
							url = fmt.Sprintf("%s/%s", webSite, part[firstQuote+1:secondQuote+1])

							fmt.Println(url)

							s.getImageFromWWW(url, imageFile)

							parts[i] = fmt.Sprintf("src=\"%s\"", imageFile.Name()[1:])
						}
					}
				}
			}
		
			line = strings.Join(parts, " ")

		}

		fmt.Fprintln(file, line)
	}
}

func (s *Server)httpStartHandler(w http.ResponseWriter, r *http.Request) {
	url := r.FormValue("webSite")

	if len(url) == 0 {
		fmt.Println("Пустой url")
		return
	}

	fmt.Printf("Сформированы ссылка для сайта: %s\n", url)

	s.getUrlFromWWW(url, url)

	http.ServeFile(w, r, "out.html")
}

func main() {
	if len(os.Args) != 3 {
		fmt.Printf("Usage: %s <ip> <HTTP_port>", os.Args[0])
		return
	}

	ip := os.Args[1]
	HTTPport := os.Args[2]

	server := Server{
		ip: ip, 
		HTTPport: HTTPport, 
		countCSSfiles: 0,
		countImagefiles: 0,
	}

	// Маршрут для статических файлов (CSS, JS и т.д.)
    fs := http.FileServer(http.Dir("static"))
    http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", server.httpStartHandler)
	fmt.Printf("HTTP Server is listening on %s:%s\n", server.ip, server.HTTPport)
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", server.ip, server.HTTPport), nil))
}