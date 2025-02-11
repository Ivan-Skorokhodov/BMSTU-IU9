package main

import (
	"fmt"
	"net/smtp"
	"os"
	"time"
)

func main() {
	from := "smtplabaskor@mail.ru"
	password := "6tutnurgQR3XyWbrhpha"

	toList := []string{"dreikonline5554545@gmail.com"}

	host := "smtp.mail.ru"
	port := "587"

	auth := smtp.PlainAuth("", from, password, host)

	for i := 0; i < len(toList); i++ {
		fmt.Println("send mail to: " + toList[i])

		subject := "Рецепт пельменей от Скороходового Ивана"

		msg := "<p>Уважаемый " + toList[i] + ",</p>\n" +
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

		msgRes := fmt.Sprintf("To: %s\r\nFrom: %s\r\nSubject: %s\r\n", toList[i], from, subject)
		msgRes += "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\r\n"
		msgRes += fmt.Sprintf("\r\n%s\r\n", msg)

		err := smtp.SendMail(host+":"+port, auth, from, []string{toList[i]}, []byte(msgRes))

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		time.Sleep(2)
	}

	fmt.Println("Successfully sent mail to all users in toList")
}
