package main

import (
	"fmt"
	"math/big"
)

/*
Разработать программу, которая перемножает, делит, складывает, вычитает две числовых переменных a,b,
значение которых > 2^20.
*/

func bigSum(a, b *big.Int) *big.Int {
	return new(big.Int).Add(a, b)
}

func bigMul(a, b *big.Int) *big.Int {
	return new(big.Int).Mul(a, b)
}

func bigDiv(a, b *big.Int) *big.Int {
	return new(big.Int).Div(a, b)
}

func bigSub(a, b *big.Int) *big.Int {
	return new(big.Int).Sub(a, b)
}

func main() {
	// задачем большие числа
	a, _ := new(big.Int).SetString("9923372036854775807978", 10)
	b, _ := new(big.Int).SetString("9923372036854775807000", 10)

	fmt.Println(bigSum(a, b))
	fmt.Println(bigMul(a, b))
	fmt.Println(bigDiv(a, b))
	fmt.Println(bigSub(a, b))
}
