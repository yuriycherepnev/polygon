package main

func main() {

}

/* через for range */
/* нет доп расходов памяти */
func charCount(str string) int {
	count := 0
	for range str {
		count++
	}
	return count
}

/* через слайс рун */
/* дополнительный расход памяти */
/*
func charCount(str string) int {
	return len([]rune(str))
}
*/
