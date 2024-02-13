package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
	=== Задача №5 ===

	Разработать программу, которая будет последовательно отправлять
	значения в канал, а с другой стороны канала — читать. По истечению
	N секунд программа должна завершаться.

*/

func main() {
	// Объявление переменной N
	var N int
	// Вывод в консоль вопроса
	fmt.Print("How many seconds will the program last? ")
	// Ввод переменной N из консоли
	fmt.Scan(&N)

	// Инициализация таймера для завершения работы программы
	timer := time.After(time.Duration(N) * time.Second)

	// Инициализация главного канала
	mainChan := make(chan int, 1)

	// Запуск бесконечного цикла
	for {
		select {
		// В случае завершения таймера
		case <-timer:
			close(mainChan)
			fmt.Println("It's time")
			return
		// В случае возможности прочитать значение из канала
		case val := <-mainChan:
			fmt.Println("read:", val)
		// В случае возможности записать значение в канала
		case mainChan <- rand.Intn(10):
			fmt.Println("write")
		}
	}
}
