package main

import "fmt"

/*
Дана последовательность температурных колебаний: -25.4, -27.0 13.0, 19.0, 15.5, 24.5, -21.0, 32.5.
Объединить данные значения в группы с шагом в 10 градусов. Последовательность в подмножноствах не важна.
*/

func main() {
	temptGroups := make(map[int][]float64)
	tempts := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	for _, t := range tempts {
		tempGroup := int(t/10) * 10
		temptGroups[tempGroup] = append(temptGroups[tempGroup], t)
	}
	fmt.Println(temptGroups)
}
