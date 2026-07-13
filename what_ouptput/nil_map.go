package main

import "fmt"

func main() {
	var m map[string]int

	m["A"] = 1
	m["B"] = 2
	m["C"] = 3

	fmt.Println(m)
}

/*
здесь будет panic
nil map не работает на запись, но работает на чтение
Исправить можно через make:
m := make(map[string]int)
*/
