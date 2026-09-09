package main

import (
	"database/sql"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
)

// ТЗ: мы хотим собирать информацию по курсам валют в разных банках.
// Требуется написать программу, которая каждую минуту будет отправлять запрос в банк и
// получать курс нескольких валют и сохранять результат в БД.
// Банков может быть несколько.

func main() {
	if len(os.Args) == 2 {
		cmdName := os.Args[1]
		if cmdName == "help" {
			fmt.Println("Usage is './currency update'")
		} else if cmdName == "update" {
			urlsBank := []struct {
				bankName string
				curFrom  string
				curTo    string
				url      string
			}{
				{
					bankName: "Bank 1",
					curFrom:  "RUB",
					curTo:    "USD",
					url:      "http://bank.example.com/rates/rub-usd",
				},
				{
					bankName: "Bank 2",
					curFrom:  "RUB",
					curTo:    "USD",
					url:      "http://bank2.example.com/rates?currFrom=RUR&currTo=USD",
				},
			}
			clientBank := &http.Client{}
			for _, url := range urlsBank {
				req, _ := http.NewRequest(http.MethodGet, url.url, nil)
				if url.bankName == "Bank 2" {
					req.Header.Add("Authorization", "auth_token=\"XXXXXXX\"")
				}
				resp, err := clientBank.Do(req)
				if err != nil {
					panic(err)
				}
				defer resp.Body.Close()
				body, _ := io.ReadAll(resp.Body)
				strBody := string(body)
				if url.bankName == "Bank 1" {
					strBody = strings.ReplaceAll(strBody, ",", ".") // Заменяем для Банка 1 запятую на точку
				}
				value, err := strconv.ParseFloat(strBody, 64)
				if err != nil {
					panic(err)
				}
				err = updateCurrency(url.bankName, url.curFrom, url.curFrom, value)
				if err != nil {
					panic(err)
				}
			}
		}
	} else {
		fmt.Println("Usage is './currency update'")
	}
}

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "<password>"
	dbname   = "<dbname>"
)

func updateCurrency(bank, from, to string, value float64) error {
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)
	defer db.Close()
	err = db.Ping()
	CheckError(err)
	fmt.Println("Connected!")
	insertStmt := fmt.Sprintf(`insert into currency_rates ("bank", "from", "to", "value") values('%s', '%s', '%s', '%.2f')`, bank, from, to, value)
	_, err = db.Exec(insertStmt)
	return err
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

/*
1. Инициализация параметров в args в отдельной функции
2. if cmdName == "help" переписать на switch case
3. обновление базы данных в отдельной функции
4. urlsBank - захардкожены, необходимо вынести в бд
5. прописать структуру для банка
6. запросы в банки можно сделать асинхронно в горутинах (worker pool)
7. url в for переименовать в bank
8. http запрос к банку в отдельную функцию
9. нет проверки ошибки в запросе к банку req, _ := http.NewRequest
10. if bank.bankName == "Bank 2" это костыль
заголовки надо вынести в мапу в структуру bankStorage
11. bank.bankName == "Bank 1" это также костыль
необходимо вынести в мапу rules внутри bankRepository
12. Заменяем для Банка 1 запятую на точку - лишний ненужный комментарий
13. обработка ошибок должна быть на логировании, а не через панику
- можно написать кастомные типы ошибок
- сделать обертки и т.д.
14. дабавить проверку статуса ответа
15. defer обьявлен внутри for, но отработает при выходе из main
- обработку банка надо вынести в отдельную функцию
16. баг, дважды написан curFrom
17. вставку в бд можно сделать батчем + транзакция, а не обновлять построчно
18. конфиги для доступа к бд должны лежать в переменных окружения
либо подгружаться из кого то хранилища паролей (vault)
19. подключение и инициализация бд должны быть вынесены
в отдельную сущность
- в идеале какой то singleton для доступа к бд
20. fmt.Println("Connected!") это пишется в логи
21. вставка в бд должна проходить валидацию для защиты от SQL иньекций
- в идеале каждый параметр должен проходить отдельную валидацию на тип и прочие проверки
*/
