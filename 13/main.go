package main

import "fmt"

/*
Поменять местами два числа без создания временной переменной.
*/

func main() {
	a := 2
	b := 5
	a, b = b, a
	fmt.Printf("a=%v, b=%v\n", a, b)
}
