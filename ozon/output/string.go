package main

func main() {
	s := "test"

	println(s[0]) // что выведет?

	b := []rune(s)
	b[0] = 'R'
	s = string(b)

	println(s)
}
