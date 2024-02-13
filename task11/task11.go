package main

import (
	"fmt"
	"strings"
)

/*
	=== Задача №11 ===

	Реализовать пересечение двух неупорядоченных множеств.

*/

func main() {
	arr1 := []int{1, 3, 5, 7, 9}
	arr2 := []int{1, 2, 3, 4, 5}
	m := make(map[int]struct{})

	for _, v := range arr1 {
		m[v] = struct{}{}
	}

	res := make([]int, 0)
	for _, v := range arr2 {
		_, ok := m[v]
		if ok {
			res = append(res, v)
		}
	}

	fmt.Println("set1: ", strings.Trim(fmt.Sprint(arr1), "[]"))
	fmt.Println("set2: ", strings.Trim(fmt.Sprint(arr2), "[]"))
	fmt.Println("result: ", strings.Trim(fmt.Sprint(res), "[]"))
}
