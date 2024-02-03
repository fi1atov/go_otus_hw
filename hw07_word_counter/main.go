package main

import (
	"fmt"
	"strings"
)

func split(r rune) bool {
	return r == '.' || r == ';' || r == ',' || r == ' '
}

func countWords(row string) (res map[string]int) {
	res = map[string]int{}
	for _, word := range strings.FieldsFunc(row, split) {
		res[word]++
	}
	return
}

func main() {
	fmt.Println(countWords("ht hl;hl ht,gf ht"))
}
