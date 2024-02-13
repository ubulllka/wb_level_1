package main

import "fmt"

/*
	=== Задача №13 ===

	Поменять местами два числа без создания временной переменной.

*/

func main() {
	a := 22
	b := 33
	fmt.Print("a = ", a, ", b = ", b, "\n")
	
	a = a + b
	b = a - b
	a = a - b
	fmt.Print("a = ", a, ", b = ", b, "\n")
	
	a, b = b, a
	fmt.Print("a = ", a, ", b = ", b, "\n")
}
