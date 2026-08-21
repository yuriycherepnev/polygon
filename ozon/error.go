package main

import (
	"database/sql"
	"errors"
)

func main() {
	err := handle()
	println(err)
	errors.Is(err, sql.ErrNoRows)
}

func handle() error {
	return &customError{}
}

type customError struct{}

func (e customError) Error() string {
	return "custom error"
}
