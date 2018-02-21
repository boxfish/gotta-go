package main

import (
	"bytes"
	"errors"
	"fmt"
	"log"
	"strings"
	"unicode/utf8"
)

func parseSymbol(char rune, base int) (value int, err error) {
	switch {
	case char >= '0' && char <= '9':
		// from 0 - 9
		value = int(char - '0')
	case char >= 'A' && char <= 'Z':
		// from A - Z
		value = 10 + int(char-'A')
	case char >= 'a' && char <= 'z':
		// from a - z
		value = 10 + int(char-'a')
	default:
		return value, errors.New("unrecognized symbol")
	}
	if value >= base {
		return value, errors.New("out of range")
	}
	return value, nil
}

func toSymbol(value int) (char rune, err error) {
	switch {
	case value >= 0 && value <= 9:
		char = rune(value) + '0'
	case value >= 10 && value <= 10+'Z'-'A':
		char = rune(value-10) + 'A'
	default:
		return char, errors.New("out of vocabulary")
	}
	return char, nil
}

func parseNumber(text string, base int) (result int, err error) {
	for _, char := range text {
		v, err := parseSymbol(char, base)
		if err != nil {
			return result, err
		}
		result = result*base + v
	}
	return result, nil
}

func toString(value int, base int) (result string, err error) {
	var buffer bytes.Buffer
	for ; value > 0; value /= base {
		remainder := value % base
		char, err := toSymbol(remainder)
		if err != nil {
			return result, err
		}
		buffer.WriteRune(char)
	}
	reversedBytes := buffer.Bytes()
	var resultBuilder strings.Builder
	for {
		r, size := utf8.DecodeLastRune(reversedBytes)
		if r == utf8.RuneError {
			break
		}
		resultBuilder.WriteRune(r)
		reversedBytes = reversedBytes[:len(reversedBytes)-size]
	}
	return resultBuilder.String(), nil
}

func main() {
	fromText := "eeff00"
	fromBase := 16
	toBase := 10
	n, err := parseNumber(fromText, fromBase)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("parse number %s from Base %v: %v\n", fromText, fromBase, n)
	toText, err := toString(n, toBase)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("convert %v to Base %v: %s\n", n, toBase, toText)
}
