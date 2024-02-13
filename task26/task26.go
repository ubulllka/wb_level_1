package main

import "fmt"

/*
	=== Задача №26 ===

	Разработать программу, которая проверяет, что все символы в строке
	уникальные (true — если уникальные, false etc). Функция проверки
	должна быть регистронезависимой.

	Например:
	abcd — true
	abCdefAaf — false
	aabcd — false

*/


func isUniq(str string) bool {
	mp := make(map[rune]struct{}, len(str))
	for _, v := range str {
		_, ok := mp[v]
		if ok {
			return false
		}
		mp[v] = struct{}{}
	}
	return true
}

func main() {
	str1 := "abcd"
	fmt.Println(str1, isUniq(str1))
	str2 := "abCdefAaf"
	fmt.Println(str2, isUniq(str2))
	str3 := "aabcd"
	fmt.Println(str3, isUniq(str3))
}
