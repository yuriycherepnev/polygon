package main

import "fmt"

func main() {
	runeVar := digitRune(7)
	fmt.Println(runeVar)

	str := runeString(runeVar)
	fmt.Println(str)

	digit := runeDigit('7')
	fmt.Println(digit)
}

func runeDigit(s rune) int {
	return int(s - '0')
}

func digitRune(s int) rune {
	return rune('0' + s)
}

func runeString(runeVar rune) string {
	return string(runeVar)
}

func stringRune(str string) []rune {
	return []rune(str)
}
