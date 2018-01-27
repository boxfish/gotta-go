// reverse the characters of a []byte slice that represents a UTF-8-encoded string, in place. Can you do it without allocating new memory?

package main

import (
	"fmt"
	"unicode/utf8"
)

// reverse the bytes for each
func reverse(s []byte) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

// given a byte slice representing a utf-8 string, reverse the runes
func reverseUTF8Runes(s []byte) {
	// pass 1: reverse each rune in place
	for i := 0; i < len(s); {
		for j := i + 1; j <= len(s); j++ {
			if j == len(s) || utf8.RuneStart(s[j]) {
				// find a rune, reverse it
				reverse(s[i:j])
				i = j
				break
			}
		}
	}

	// pass 2: reverse the whole slice
	reverse(s)
}

func (s []int) len() {
	return len(s)
}

func main() {
	var s1 = "hello, 世界!"
	fmt.Println(s1)
	b1 := []byte(s1)
	reverseUTF8Runes(b1)
	fmt.Println(string(b1))
}
