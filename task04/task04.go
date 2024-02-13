package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

/*
	=== Задача №4 ===

	Реализовать постоянную запись данных в канал (главный поток). Реализовать
	набор из N воркеров, которые читают произвольные данные из канала
	и выводят в stdout. Необходима возможность выбора количества
	воркеров при старте.

	Программа должна завершаться по нажатию Ctrl+C.
	Выбрать и обосновать способ завершения работы
	всех воркеров.

*/

func main() {
	// Объявление переменной N
	var N int
	// Вывод в консоль вопроса
	fmt.Print("How many worker do you need? ")
	// Ввод переменной N из консоли
	fmt.Scan(&N)

	// Инициализация канала для отслеживания сигналов ОС
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT)

	// Инициализация главного канала
	mainChan := make(chan int)

	// Инициализация канала для завершения работы горутин-воркеров
	quit := make(chan bool)

	// Объявление переменной типа sync.WaitGroup
	var wg sync.WaitGroup

	// Увеличения счетчика горутин в переменной wg
	wg.Add(1)
	// Запуск горутины
	go func(main chan<- int, q chan<- bool) {
		// Запуск бесконечного цикла
		for {
			select {
			// В случае сигнала от системы
			case <-sigChan:
				// Отправление сигнала о завершении работы
				q <- true
				wg.Done()
				fmt.Println("Press ctrl + c")
				return
			default:
				// Отправление в канал целого числа
				mainChan <- rand.Intn(10)
			}

		}
	}(mainChan, quit)

	// Цикл по воркерам
	for i := 0; i < N; i++ {
		// Увеличения счетчика горутин в переменной wg
		wg.Add(1)
		// Запуск горутины воркера
		go func(main <-chan int, i int) {
			// Уменьшение счетчика горутин перед выходом из функции
			defer wg.Done()

			select {
			// В случае сигнала о завершении работы
			case <-quit:
				return
			default:
				fmt.Printf("Read worker %d: %d\n", i, <-main)
			}
		}(mainChan, i)
		// Небольшая пауза для медленной работы воркеров
		time.Sleep(500 * time.Nanosecond)
	}

	// Ожидание завершания работы всех горутин
	wg.Wait()
}
