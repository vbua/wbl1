package main

import "fmt"

/*
Дана переменная int64. Разработать программу которая устанавливает i-й бит в 1 или 0.
*/

// читал здесь - https://medium.com/learning-the-go-programming-language/bit-hacking-with-go-e0acee258827

func setBit(i int, number int64, k int) int64 {
	mask := int64(1) << (k - 1)
	fmt.Printf("mask: %064b\n", mask)
	if i == 1 {
		/*
			биты сравниваются между собой и если хотя бы один бит из пары равен 1, то вернем 1, иначе 0
		*/
		number |= mask
	} else if i == 0 {
		/*
			AND_NOT(a, b) = AND(a, NOT(b))
			то есть AND_NOT(a, 1) = 0; так можно перевести определенные биты в 0
		*/
		number &^= mask
	}

	return number
}

func main() {
	var number int64 = 255
	fmt.Printf("%064b\n", number)
	fmt.Printf("%064b\n", setBit(0, number, 63))
}
