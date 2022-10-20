package main

import (
	"fmt"
)

/*
Дана структура Human (с произвольным набором полей и методов). Реализовать встраивание методов в структуре Action
от родительской структуры Human (аналог наследования).
*/

type Human struct {
	name string
	age  int
}

// Создаю методы структуры Human
func (h *Human) Eat() {
	fmt.Printf("%v eats\n", h.name)
}

func (h *Human) Sleep() {
	fmt.Printf("%v sleeps\n", h.name)
}

/*
Встраиваю структуру Human в структуру Action. В данном случае Human родительская структура.
Action "наследует" все методы структуры Human.
*/
type Action struct {
	Human
}

func main() {
	action := Action{
		Human: Human{
			name: "John",
			age:  25,
		},
	}
	// Вызывается метод Eat из структуры Human
	action.Eat()
	// Вызывается метод Sleep из структуры Human
	action.Sleep()
}
