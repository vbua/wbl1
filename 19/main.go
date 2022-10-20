package main

import "fmt"

/*
Разработать программу, которая переворачивает подаваемую на ход строку (например: «главрыба — абырвалг»).
Символы могут быть unicode.
*/

func reverseString(text string) string {
	runedText := []rune(text)
	for i, l := 0, len(runedText)-1; i < l; i, l = i+1, l-1 {
		runedText[i], runedText[l] = runedText[l], runedText[i]
	}
	return string(runedText)
}

func main() {
	fmt.Println(reverseString("главрыба"))
}
