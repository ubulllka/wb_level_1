package main

import (
	"fmt"
	"strconv"
)

/*
	=== Задача №8 ===

	Дана переменная int64. Разработать программу которая устанавливает i-й бит в 1 или 0.

*/

func main() {
	var n, i, k int64

	fmt.Print("Input random number: ")
	fmt.Scan(&n)

	str := strconv.FormatInt(n, 2)
	fmt.Printf("number: \"%d\" in 10 and \"%s\" in 2\n", n, str)

	fmt.Printf("Input digit of the number (<= %d): ", len(str))
	fmt.Scan(&i)

	fmt.Print("Input 0 or 1: ")
	fmt.Scan(&k)

	var newN int64
	if k == 0 {
		newN = n & ^(1 << i)
	} else {
		newN = n | (1 << i)
	}
	fmt.Printf("number: \"%d\" in 10 and \"%s\" in 2\n", newN, strconv.FormatInt(newN, 2))
}
