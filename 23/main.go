package main

import "fmt"

/*
Удалить i-ый элемент из слайса.
*/

func removeElFromSlice(slice []int, s int) []int {
	return append(slice[:s], slice[s+1:]...)
}

func main() {
	nums := []int{5, 8, 10, 45, 11, 22, 34}
	fmt.Println(removeElFromSlice(nums, 3))
}
