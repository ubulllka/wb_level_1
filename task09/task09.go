package main

import (
	"fmt"
	"sync"
)

/*
	=== Задача №9 ===

	Разработать конвейер чисел. Даны два канала: в первый пишутся
	числа (x) из массива, во второй — результат операции x*2,
	после чего данные из второго канала должны выводиться в stdout.

*/

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	firstChan := make(chan int)
	secondChan := make(chan int)

	var wg sync.WaitGroup

	wg.Add(3)

	// Запись в первый канал
	go func(arr []int, f chan<- int) {
		defer wg.Done()
		defer close(f)

		for _, v := range arr {
			f <- v
		}

	}(arr, firstChan)

	// Запись во второй канал
	go func(f <-chan int, s chan<- int) {
		defer wg.Done()
		defer close(s)

		for elem := range f {
			s <- elem * 2
		}

	}(firstChan, secondChan)

	// Вывод в консоль
	go func(s chan int) {
		defer wg.Done()

		for elem := range s {
			fmt.Println(elem)
		}

	}(secondChan)

	wg.Wait()
}
