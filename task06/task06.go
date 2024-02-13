package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

/*
	=== Задача №6 ===

	Реализовать все возможные способы остановки выполнения горутины.

*/

// Остановка с помощью канала
func first() {
	quit := make(chan struct{})

	var wg sync.WaitGroup

	wg.Add(1)

	go func(q <-chan struct{}) {
		for {
			select {
			case <-q:
				defer wg.Done()
				fmt.Println("First goroutine stopped...")
				return
			default:
				fmt.Println("First goroutine work!")
				time.Sleep(time.Second)
			}
		}
	}(quit)
	time.Sleep(3 * time.Second)
	quit <- struct{}{}
	wg.Wait()
}

// Остановка через указатель
func second() {
	quit := false

	var wg sync.WaitGroup

	wg.Add(1)

	go func(q *bool) {
		for {
			if *q {
				defer wg.Done()
				fmt.Println("Second goroutine stopped...")
				return
			} else {
				fmt.Println("Second goroutine work!")
				time.Sleep(time.Second)
			}
		}
	}(&quit)
	time.Sleep(3 * time.Second)
	quit = true
	wg.Wait()
}

// Остановка через закрытия контекста
func third() {
	quit := make(chan struct{})

	ctx, cancel := context.WithCancel(context.Background())

	var wg sync.WaitGroup

	wg.Add(1)

	go func(c *context.Context, q chan struct{}) {
		for {
			select {
			case <-ctx.Done():
				defer wg.Done()
				fmt.Println("Third goroutine stopped...")
				return
			default:
				fmt.Println("Third goroutine work!")
				time.Sleep(time.Second)
			}
		}
	}(&ctx, quit)
	time.Sleep(3 * time.Second)
	cancel()
	wg.Wait()
}

func main() {
	first()
	second()
	third()
}
