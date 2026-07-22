/*
Здесь проблема в том, что в слайс копируется ссылка на массив
и при дальнейшей работе, garbage collector не может его удалить,
чтоб это исправить нужно использовать функцию copy

Функция copy в Go копирует элементы из одного слайса (или массива) в другой.
Копируется минимум из двух длин.

В отличие от append, copy никогда не увеличивает длину или емкость слайса.
*/

package main

import "fmt"

func getBytes(start, end int) []byte {
	arr := [999]byte{}

	slice := arr[start:end]

	/*
		slice2 := make([]byte, end-start)
		copy(slice2, arr[start:end])
	*/

	return slice
}

func main() {
	s := getBytes(10, 20)

	fmt.Println(s)
	fmt.Println(cap(s))
}
