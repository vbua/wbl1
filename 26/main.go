package main

import (
	"fmt"
	"strings"
)

/*
Разработать программу, которая проверяет, что все символы в строке уникальные (true — если уникальные, false etc).
Функция проверки должна быть регистронезависимой.

Например:
abcd — true
abCdefAaf — false
aabcd — false
*/

func AreUniqueSymbols(str string) bool {
	str = strings.ToLower(str)
	hash := make(map[rune]int)
	for _, v := range str {
		if _, ok := hash[v]; ok {
			return false
		}
		hash[v]++
	}
	return true
}

func main() {
	fmt.Println(AreUniqueSymbols("abcd"))
	fmt.Println(AreUniqueSymbols("abCdefAaf"))
	fmt.Println(AreUniqueSymbols("aabcd"))
}
