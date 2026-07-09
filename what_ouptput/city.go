package what_ouptput

import "fmt"

type Address struct {
	city   string
	street string
	house  int
}

// не отработает, т.к. a это просто копия ссылка.
// чтоб поменять значение, нужно сделать либо мутацию обьекта, a.city
// либо дереференс и замену обьекта
func (a *Address) setCity(city string) {
	a = &Address{
		city: city,
	}
}

func (a *Address) setStreet(street string) {
	a.street = street
}

// тоже самое, здесь в addr лежит копия ссылки
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
