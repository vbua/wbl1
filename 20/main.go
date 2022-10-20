package main

import (
	"fmt"
	"strings"
)

/*
Разработать программу, которая переворачивает слова в строке.
Пример: «snow dog sun — sun dog snow».
*/

func reverseWordsInString(text string) string {
	words := strings.Split(text, " ")
	for i, l := 0, len(words)-1; i < l; i, l = i+1, l-1 {
		words[i], words[l] = words[l], words[i]
	}
	return strings.Join(words, " ")
}

func main() {
	fmt.Println(reverseWordsInString("snow dog sun читать есть спать"))
}
