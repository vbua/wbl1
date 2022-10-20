package main

import (
	"fmt"
)

/*
Реализовать бинарный поиск встроенными методами языка.
*/

func binarySearch(nums []int, item int) int {
	low, high := 0, len(nums)-1
	for low <= high {
		mid := (low + high) / 2
		guess := nums[mid]
		switch {
		case guess == item:
			return mid
		case guess > item:
			high = mid - 1
		default:
			low = mid + 1
		}
	}
	return -1
}

func main() {
	fmt.Println(binarySearch([]int{1, 2, 3, 4, 7, 10, 15, 22, 33}, 2))
}
