package main

import (
	"fmt"
	"math/big"
)

/*
	=== Задача №22 ===

	Разработать программу, которая перемножает, делит, складывает, вычитает
	две числовых переменных a,b, значение которых > 2^20.

*/

func main() {
	a := big.NewFloat(9876543212345679)
	b := big.NewFloat(7654322345678456)

	sum := new(big.Float)
	fmt.Println(a, "+", b, "=", sum.Add(a, b))

	sub := new(big.Float)
	fmt.Println(a, "-", b, "=", sub.Sub(a, b))

	mul := new(big.Float)
	fmt.Println(a, "*", b, "=", mul.Mul(a, b))
	
	div := new(big.Float)
	fmt.Println(a, "/", b, "=", div.Quo(a, b))
}
