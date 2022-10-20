package main

import (
	"fmt"
	"math/rand"
)

/*
Реализовать быструю сортировку массива (quicksort) встроенными методами языка.
*/

/*
суть в том, что мы выбираем рандомный базовый элемент и кладем слева от него числа меньше него, а справа больше него.
и так рекурсивно по правой и левой части.
*/

func quickSort(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}

	left, right := 0, len(nums)-1

	// рандомно выбираем индекс базового элемента
	pivotIndex := rand.Int() % len(nums)

	// базовый элемент отправляем вправо
	nums[pivotIndex], nums[right] = nums[right], nums[pivotIndex]

	// все элементы, которые меньше базового элемента, складируем слева
	for i := range nums {
		if nums[i] < nums[right] {
			nums[i], nums[left] = nums[left], nums[i]
			left++
		}
	}

	// базовый элемент ставим после всех элементов, которые меньше его
	nums[left], nums[right] = nums[right], nums[left]

	/*
	 в итоге у нас получается слайс, в котором слева элементы меньше базового элемента, а справа - больше.
	*/

	// рекурсивно повторяем тоже самое с левой и правой частью
	quickSort(nums[:left])
	quickSort(nums[left+1:])
	return nums
}

func main() {
	fmt.Println(quickSort([]int{1, 4, 6, 7, 10, 5, 2, 3, 3}))
	fmt.Println(quickSort([]int{15, 7}))
}
