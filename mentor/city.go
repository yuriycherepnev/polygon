package main

import "fmt"

type Address struct {
	city   string
	street string
	house  int
}

// не отработает, здесь в a копия ссылки, а не сама структура Address
// чтоб поменять значение, нужно сделать либо мутацию обьекта через a.city
// либо дереференс и замену обьекта
func (a *Address) setCity(city string) {
	a = &Address{
		city: city,
	}
}

// не отработает, т.к. передана копия структуры
func (a Address) setStreet(street string) {
	a.street = street
}

// тоже самое, здесь addr копия ссылки, а не сама структура
func setHouse(addr *Address, house int) {
	addr = &Address{
		house: house,
	}
}

func main() {
	addr := Address{
		city:   "New York",
		street: "Broadway",
		house:  10,
	}

	addr.setCity("London")
	addr.setStreet("Piccadilly")
	setHouse(&addr, 5)

	fmt.Println(addr)
}
