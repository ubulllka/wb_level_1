package main

import (
	"fmt"
	"strings"
)

/*
	=== Задача №20 ===

	Разработать программу, которая переворачивает слова в строке.

	Пример: «snow dog sun — sun dog snow».

*/
func reverseStr(str string) string {
	var s strings.Builder
	var temp strings.Builder
	temp.WriteString(" ")
	for _, v := range str {
		if v == ' ' {
			temp.WriteString(s.String())
			s = temp
			temp.Reset()
			temp.WriteString(" ")
		} else {
			temp.WriteRune(v)
		}
	}
	temp.WriteString(s.String())
	s = temp
	return s.String()[1:]
}

func main() {
	str := "snow dog sun"
	res := reverseStr(str)

	fmt.Println(res, res == "sun dog snow")
}
