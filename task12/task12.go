package main

import (
	"fmt"
	"strings"
)

/*
	=== Задача №12 ===

	Имеется последовательность строк - (cat, cat, dog, cat, tree) создать для нее собственное множество.

*/

func main() {
	arr := []string{"cat", "cat", "dog", "cat", "tree"}

	m := make(map[string]struct{})

	for _, v := range arr {
		m[v] = struct{}{}
	}

	res := make([]string, 0, len(m))
	for k, _ := range m {
		res = append(res, k)
	}

	fmt.Println(strings.Trim(fmt.Sprint(res), "[]"))
}
