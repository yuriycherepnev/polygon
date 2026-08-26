package main

import (
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
	runePwd := []rune(password)
	if len(runePwd) == length {
		if checkPassword(password, hash) {
			return password
		}
		return ""
	}
	for _, char := range alphabet {
		result := generatePassword(hash, password+string(char), length)

		if result != "" {
			return result
		}
	}
	return ""
}

func checkPassword(password string, hash []byte) bool {
	hashPass := hashPassword(password)
	return string(hashPass) == string(hash)
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
