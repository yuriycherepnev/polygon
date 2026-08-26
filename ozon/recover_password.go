package main

import (
	"bytes"
	"crypto/md5"
	"fmt"
)

var alphabet = []rune{'a', 'b', 'c', 'd', '1', '2', '3'}

func main() {
	TestRecoverPassword()
}

func RecoverPassword(hash []byte) string {
	length := 1
	for {
		result := generatePassword(hash, "", length)
		if result != "" {
			return result
		}
		length++
	}
}

func generatePassword(hash []byte, password string, length int) string {
	if len(password) == length {
		fmt.Println(password)
		if checkPassword(password, hash) {
			return password
		}
		return ""
	}
	for _, ch := range alphabet {
		result := generatePassword(hash, password+string(ch), length)

		if result != "" {
			return result
		}
	}
	return ""
}

func checkPassword(password string, h []byte) bool {
	hash := hashPassword(password)
	return bytes.Equal(hash, h)
}

func TestRecoverPassword() {
	for _, exp := range []string{
		"2",
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
