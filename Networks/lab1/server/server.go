package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"math/big"
	"net"
	"strconv"

	log "github.com/mgutz/logxi/v1"
)

// Client - состояние клиента.
type Client struct {
	logger   log.Logger    // Объект для печати логов
	conn     *net.TCPConn  // Объект TCP-соединения
	enc      *json.Encoder // Объект для кодирования и отправки сообщений
	sum      *big.Rat      // Текущая сумма полученных от клиента дробей
	list_num []int         // Список чисел
}

// NewClient - конструктор клиента, принимает в качестве параметра
// объект TCP-соединения.
func NewClient(conn *net.TCPConn) *Client {
	return &Client{
		logger: log.New(fmt.Sprintf("client %s", conn.RemoteAddr().String())),
		conn:   conn,
		enc:    json.NewEncoder(conn),
		list_num:  []int{},
	}
}

// serve - метод, в котором реализован цикл взаимодействия с клиентом.
// Подразумевается, что метод serve будет вызаваться в отдельной go-программе.
func (client *Client) serve() {
	defer client.conn.Close()
	decoder := json.NewDecoder(client.conn)
	for {
		var req Request
		if err := decoder.Decode(&req); err != nil {
			client.logger.Error("cannot decode message", "reason", err)
			break
		} else {
			client.logger.Info("received command", "command", req.Command)
			if client.handleRequest(&req) {
				client.logger.Info("shutting down connection")
				break
			}
		}
	}
}

// handleRequest - метод обработки запроса от клиента. Он возвращает true,
// если клиент передал команду "quit" и хочет завершить общение.
func (client *Client) handleRequest(req *Request) bool {
	switch req.Command {
	case "quit":
		client.respond("ok", nil)
		return true
	case "add":
		errorMsg := ""
		if req.Data == nil {
			errorMsg = "data field is absent"
		} else {
			var frac Number
			if err := json.Unmarshal(*req.Data, &frac); err != nil {
				errorMsg = "malformed data field"
			} else {
				n := frac.Num
				client.logger.Info("performing addition number to list", "value", n)
				number, _ := strconv.Atoi(n)
				client.list_num = append(client.list_num, number)
			}
		}
		if errorMsg == "" {
			client.respond("ok", nil)
		} else {
			client.logger.Error("addition failed", "reason", errorMsg)
			client.respond("failed", errorMsg)
		}
	case "avg":
		if len(client.list_num) == 0 {
			client.logger.Error("calculation failed", "reason", "zero len list")
			client.respond("failed", "zero len list")
		} else {

			errorMsg := ""
			if req.Data == nil {
				errorMsg = "data field is absent"
			} else {
				var n2 TwoNumbers
				if err := json.Unmarshal(*req.Data, &n2); err != nil {
					errorMsg = "malformed data field"
				} else {
					fn := n2.FNum
					sn := n2.SNum
					client.logger.Info("good edges of list", "values", fn, sn)

					fnumber, _ := strconv.Atoi(fn)
					snumber, _ := strconv.Atoi(sn)

					mx := client.list_num[fnumber]
					for i := fnumber; i <= snumber; i++ {
						if mx < client.list_num[i] {
							mx = client.list_num[i]
						}
					}

					mx_s := strconv.Itoa(mx)

					client.respond("result", &Number{
						Num:   mx_s,
					})

				}
			}

			_ = errorMsg
			/*
			if errorMsg == "" {
				client.respond("ok", nil)
			} else {
				client.logger.Error("addition failed", "reason", errorMsg)
				client.respond("failed", errorMsg)
			}
			*/
		}
	default:
		client.logger.Error("unknown command")
		client.respond("failed", "unknown command")
	}
	return false
}

// respond - вспомогательный метод для передачи ответа с указанным статусом
// и данными. Данные могут быть пустыми (data == nil).
func (client *Client) respond(status string, data interface{}) {
	var raw json.RawMessage
	raw, _ = json.Marshal(data)
	client.enc.Encode(&Response{status, &raw})
}

func main() {
    // Работа с командной строкой, в которой может указываться необязательный ключ -addr.
	var addrStr string
	flag.StringVar(&addrStr, "addr", "185.104.251.226:6001", "specify ip address and port")
	flag.Parse()

    // Разбор адреса, строковое представление которого находится в переменной addrStr.
	if addr, err := net.ResolveTCPAddr("tcp", addrStr); err != nil {
		log.Error("address resolution failed", "address", addrStr)
	} else {
		log.Info("resolved TCP address", "address", addr.String())

        // Инициация слушания сети на заданном адресе.
		if listener, err := net.ListenTCP("tcp", addr); err != nil {
			log.Error("listening failed", "reason", err)
		} else {
            // Цикл приёма входящих соединений.
			for {
				if conn, err := listener.AcceptTCP(); err != nil {
					log.Error("cannot accept connection", "reason", err)
				} else {
					log.Info("accepted connection", "address", conn.RemoteAddr().String())

                    // Запуск go-программы для обслуживания клиентов.
					go NewClient(conn).serve()
				}
			}
		}
	}
}
