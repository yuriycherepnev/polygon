// type assertion это приведение значения из interface{} / any к конкретному типу

package main

import "fmt"

func main() {
	var value any
	value = "22"
	number := typeAssertion(value)
	fmt.Println(number)
}

func typeAssertion(value any) int {
	actual, ok := value.(int)
	if !ok {
		panic("value is not int")
	}
	return actual
}
