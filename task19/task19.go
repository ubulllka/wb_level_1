package main

import (
	"fmt"
	"strings"
)

/*
	=== Задача №19 ===

	Разработать программу, которая переворачивает подаваемую на ход
	строку (например: «главрыба — абырвалг»). Символы могут быть unicode.

*/

func reverse(str string) string {
	var s strings.Builder
	for _, v := range str {
		var temp strings.Builder 
		temp.WriteRune(v)
		temp.WriteString(s.String())
		s = temp
	}
	return s.String()
}

func main() {
	str := "главрыба"
	res := reverse(str)

	fmt.Println(res, res == "абырвалг")
}
