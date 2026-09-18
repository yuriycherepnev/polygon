package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

func WeatherForecast() int {
	time.Sleep(1 * time.Second)
	return rand.Intn(70) - 30
}

func main() {
	http.HandleFunc("/weather", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "{\"temperature\":%d}\n", WeatherForecast())
	})

	if err := http.ListenAndServe(":3333", nil); err != nil {
		panic(err)
	}
}
