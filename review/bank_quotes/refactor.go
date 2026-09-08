/*
package main

import (
	"database/sql"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

// ТЗ: мы хотим собирать информацию по курсам валют в разных банках.
// Требуется написать программу, которая каждую минуту будет отправлять запрос в банк и
// получать курс нескольких валют и сохранять результат в БД.
// Банков может быть несколько.

var (
	cmdName string
)

type bankRepository struct {
}

type bankStorage struct {
	bankName string
	curFrom  string
	curTo    string
	url      string
	headers  map[string]string
	rules    map[string]string
}

func parseArgs() {
	cmdName = os.Args[1]
}

func main() {
	parseArgs()

	db, err := connectDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	switch cmdName {
	case "help":
		fmt.Println("Usage is './currency update'")
	case "update":
		update(db)
	default:
		fmt.Println("Usage is './currency update'")
	}
}

func update(db *sql.DB) {
	repository := bankRepository{}
	urlsBank := repository.getBanks()
	clientBank := &http.Client{}
	for _, bank := range urlsBank {
		value := getValue(clientBank, bank)

		err := updateCurrency(
			db,
			bank.bankName,
			bank.curFrom,
			bank.curTo,
			value,
		)
		if err != nil {
			log.Println(err)
		}
	}
}

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "<password>"
	dbname   = "<dbname>"
)

func getValue(clientBank *http.Client, bank bankStorage) float64 {
	req, err := http.NewRequest(http.MethodGet, bank.url, nil)
	if err != nil {
		log.Println(err)
	}
	for key, value := range bank.headers {
		req.Header.Add(key, value)
	}
	resp, err := clientBank.Do(req)
	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Println(err)
	}

	body, _ := io.ReadAll(resp.Body)
	strBody := string(body)

	for oldVal, newVal := range bank.rules {
		strBody = strings.ReplaceAll(strBody, oldVal, newVal)
	}

	value, err := strconv.ParseFloat(strBody, 64)
	if err != nil {
		log.Println(err)
	}

	return value
}

func connectDB() (*sql.DB, error) {
	connection := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname,
	)

	db, err := sql.Open("postgres", connection)
	if err != nil {
		log.Println(err)
	}
	if err = db.Ping(); err != nil {
		defer db.Close()
		log.Println(err)
	}

	return db, nil
}
func updateCurrency(db *sql.DB, bank, from, to string, value float64) error {
	insertStmt := fmt.Sprintf(`insert into currency_rates ("bank", "from", "to", "value") values('%s', '%s', '%s', '%.2f')`, bank, from, to, value)
	_, err := db.Exec(insertStmt)
	return err
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

func (br *bankRepository) getBanks() []bankStorage {
	return []bankStorage{
		{
			bankName: "Bank 1",
			curFrom:  "RUB",
			curTo:    "USD",
			url:      "http://bank.example.com/rates/rub-usd",
			rules: map[string]string{
				",": ".",
			},
		},
		{
			bankName: "Bank 2",
			curFrom:  "RUB",
			curTo:    "USD",
			url:      "http://bank2.example.com/rates?currFrom=RUR&currTo=USD",
			headers: map[string]string{
				"Authorization": "auth_token=\"XXXXXXX\"",
			},
		},
	}
}
*/

