package main

import (
	"crypto/md5"
	"testing"
)

var alphabet = []rune{'a', 'b', 'c', 'd', '1', '2', '3'}

func RecoverPassword(h []byte) string {
	var password string
	var generate func([]rune)

	generate = func(current []rune) {
		sum := md5.Sum([]byte(string(current)))

		if string(sum[:]) == string(h) {
			password = string(current)
			return
		}

		if password != "" {
			return
		}

		for _, ch := range alphabet {
			generate(append(current, ch))

			if password != "" {
				return
			}
		}
	}

	generate(nil)

	return password
}
func TestRecoverPassword(t *testing.T) {
	for _, exp := range []string{
		"a",
		"12",
		"abc333d",
	} {
		t.Run(exp, func(t *testing.T) {
			act := RecoverPassword(hashPassword(exp))
			if act != exp {
				t.Error("recovered:", act, "expected:", exp)
			}
		})
	}
}

func hashPassword(in string) []byte {
	h := md5.Sum([]byte(in))
	return h[:]
}
