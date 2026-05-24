package main

func main() {
	NoSpace("ass aaa")
}

func NoSpace(word string) (result string) {
	for _, char := range word {
		if char != ' ' {
			result += string(char)
		}
	}
	return result
}
