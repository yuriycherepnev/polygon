package main

import "fmt"

type Person struct {
	name string
	age  int
}

func changeName(p *Person, name string) {
	p.name = name
}

// при передаче из функции, также как и передаче в функцию, возвращаемое значение копируется
// если переменная локальная, defer ее изменения после вычисления return не будет учитывать
// если переменная обьявлена как возвращаемый результат (p Person), defer будет учитывать ее изменения
// чтоб заработал нужно прописать p как возвращаемый результат (p Person)
func getPerson() Person {
	var p Person

	p = Person{
		name: "Bob",
		age:  52,
	}

	defer changeName(&p, "Alice")
	/*
			если убрать defer, значение начнет меняться,
			связано это с тем, что defer выполняется после вычисления return,
		    и меняет только локальную переменную,
			тогда как в return уже лежит копия результата
	*/

	return p
}

func main() {
	p := getPerson()

	fmt.Println(p)
}

/*
Главное различие:
func getPerson() Person
p — обычная локальная переменная, отдельно от возвращаемого значения.

func getPerson() (p Person)
p — именованная возвращаемая переменная, и defer может изменить итоговый результат.
*/
