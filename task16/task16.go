package main

import (
	"fmt"
	"math/rand"
	"strings"
)

/*
	=== Задача №16 ===

	Реализовать быструю сортировку массива (quicksort) встроенными методами языка.

*/

func qs(arr []int, first, last int) {
	i := first
	j := last
	x := arr[int((i+j)/2)]
	for {
		for i < last && arr[i] < x{
			i++
		}
		for j > first && arr[j] > x {
			j--
		}
		if i <= j {
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
			i++
			j--
		} else {
			break
		}
	}
	if i < last {
		qs(arr, i, last)
	}
	if j > first {
		qs(arr, first, j)
	}
}

func main() {
	n := 15
	array := make([]int, n)
	for i:=0; i<n; i++ {
		array[i] = rand.Intn(90) + 10
	}
	// array := []int{66, 32, 26, 95, 29, 98, 18, 97, 32, 66, 96, 97, 67, 72, 54}
	fmt.Println("Before:", strings.Trim(fmt.Sprint(array), "[]"))
	qs(array, 0, n-1)
	fmt.Println("After: ", strings.Trim(fmt.Sprint(array), "[]"))
}
