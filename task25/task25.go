package main

import (
	"context"
	"fmt"
	"time"
)

/*
	=== Задача №25 ===

	Реализовать собственную функцию sleep.

*/

func sleep(i time.Duration) {
	ctx, cancel := context.WithTimeout(context.Background(), i*time.Second)
	defer cancel()

	<-ctx.Done()
}

func main() {
	fmt.Println("Start")
	sleep(3)
	fmt.Print("Finish")
}
