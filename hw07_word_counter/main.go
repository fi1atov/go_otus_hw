package main

import (
	"errors"
	"fmt"
	"strings"
	"unicode/utf8"
)

func split(r rune) bool {
	return r == '.' || r == ';' || r == ',' || r == ' '
}

func countWords(row string) (res map[string]int, err error) {
	res = map[string]int{}

	if utf8.ValidString(row) {
		for _, word := range strings.FieldsFunc(row, split) {
			res[word]++
		}
		return res, nil
	}
	return nil, errors.New("невалидная строка")
}

func main() {
	res, err := countWords("ht hl;hl ht,gf ht")
	fmt.Println(res, err)
	res, err = countWords("ht hl;hl ht,gf ht")
	fmt.Println(res, err)
}
