package main

import (
	"fmt"
	"sort"
)

/*
	=== Задача №17 ===

	Реализовать бинарный поиск встроенными методами языка.

*/

func binFind(_arr []int, val int) (int, bool) {
	arr := _arr
	ln := len(arr)
	if ln == 0 {
		return 0,  false
	}
	if ln == 1 {
		if arr[0] == val {
			return arr[0], true
		} else {
			return 0, false
		}
	}

	sort.Slice(arr, func(i, j int) (bool) {
		return arr[i] < arr[j]
	})

	fl := false
	lef := 0
	rig := ln - 1
	var mid int
	for lef <= rig && !fl {
		mid = int((lef + rig) / 2)
		if arr[mid] == val {
			fl = true
			return arr[mid], true
		}
		if arr[mid] > val {
			rig = mid - 1
		} else {
			lef = mid + 1
		}
	}
	return 0, false
}

func main() {
	arr := []int{1, 5, 3, 6, 8, 2, 4}
	fmt.Println(arr)
	k, v := binFind(arr, 5)
	fmt.Println("5 in arr:", k, v)
	k, v = binFind(arr, 7)
	fmt.Println("7 in arr:", k, v)
}
