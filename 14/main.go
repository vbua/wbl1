package main

import "fmt"

/*
Разработать программу, которая в рантайме способна определить тип переменной: int, string, bool, channel
из переменной типа interface{}.
*/

func getInterfaceType(i interface{}) string {
	switch i.(type) {
	case int:
		return "int"
	case string:
		return "string"
	case bool:
		return "bool"
	case chan int, chan string, chan bool, chan interface{}:
		return "channel"
	default:
		return "unknown"
	}
}

func main() {
	var i interface{}

	ch := make(chan int)
	i = ch
	fmt.Println(getInterfaceType(i))

	i = 2
	fmt.Println(getInterfaceType(i))

	i = "test"
	fmt.Println(getInterfaceType(i))

	i = true
	fmt.Println(getInterfaceType(i))
}
