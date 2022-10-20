package main

import "fmt"

/*
Разработать конвейер чисел. Даны два канала: в первый пишутся числа (x) из массива,
во второй — результат операции x*2, после чего данные из второго канала должны выводиться в stdout.
*/

/*
конвейер - это ряд этапов, связанных между собой каналами, где каждый этап это набор go-рутин
выполняющих определенную функцию.
*/

// этап отправки чисел в канал
func numberSender(numbers []int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range numbers {
			out <- n
		}
		close(out)
	}()
	return out
}

// этап умножения каждого из чисел в канале на 2
func numberDoubler(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * 2
		}
		close(out)
	}()
	return out
}

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}

	// этап вывода результатов из канала в stdout
	for v := range numberDoubler(numberSender(numbers)) {
		fmt.Println(v)
	}
}
