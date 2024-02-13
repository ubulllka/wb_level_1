package main

import (
	"fmt"
	"strings"
)

/*
	=== Задача №15 ===

	К каким негативным последствиям может привести данный фрагмент кода, и как это исправить?
	Приведите корректный пример реализации.

	var justString string
	func someFunc() {
		v := createHugeString(1 << 10)
		justString = v[:100]
	}

	func main() {
		someFunc()
	}

*/

/*
	Ответ: Строка - это слайс байт, и если мы приравниваем, а не копируем слайс или часть слайса,
	то все равно указываем на одну и ту же область памяти. Поэтому корректный пример реализации приведен ниже.
*/

var justString string

func createHugeString(n int) string {
	arr := make([]byte, n)
	for i := 0; i < n; i++ {
		arr[i] = byte('S')
	}
	return string(arr)
}

func someFunc() {
	v := createHugeString(1 << 10)
	justString = strings.Clone(v[:100])
	fmt.Println(len(justString), justString)
}

func main() {
	someFunc()
}
