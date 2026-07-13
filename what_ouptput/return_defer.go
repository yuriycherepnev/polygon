package main

import "fmt"

type Person struct {
	name string
	age  int
}

func changeName(p *Person, name string) {
	p.name = name
}

// если переменная локальная, defer ее изменения после вычисления return не будет учитывать
// если переменная обьявлена как возвращаемый результат (p Person), defer будет учитывать ее изменения
// чтоб заработал нужно прописать p как возвращаемый результат
func getPerson() Person {
	var p Person

	defer changeName(&p, "Alice")

	p = Person{
		name: "Bob",
		age:  52,
	}

	return p
}

func main() {
	p := getPerson()

	fmt.Println(p)
}
