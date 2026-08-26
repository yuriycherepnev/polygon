package main

import (
	"crypto/md5"
	"fmt"
)

var alphabet = []rune{'a', 'b', 'c', 'd', '1', '2', '3'}

func main() {
	TestRecoverPassword()
}

func RecoverPassword(h []byte) string {
	length := 1
	for {
		result := genPassword(h, nil, length)
		length++
		if result != "" {
			return result
		}
	}
}

func genPassword(h []byte, password []rune, length int) string {
	if len(password) == length {
		hash := md5.Sum([]byte(string(password)))

		if string(hash[:]) == string(h) {
			return string(password)
		}

		return ""
	}

	for _, ch := range alphabet {
		password = append(password, ch)
		result := genPassword(h, password, length)
		if result != "" {
			return result
		}

		password = password[:len(password)-1]
	}

	return ""
}

func TestRecoverPassword() {
	for _, exp := range []string{
		"a",
		"12",
		"abc333d",
	} {
		act := RecoverPassword(hashPassword(exp))

		if act != exp {
			fmt.Println(hashPassword(exp))
			fmt.Println("wrong password")
		} else {
			fmt.Println("OK:", exp)
		}
	}
}

func hashPassword(in string) []byte {
	h := md5.Sum([]byte(in))
	return h[:]
}
