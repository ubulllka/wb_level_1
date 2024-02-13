package main

import (
	"fmt"
	"sync"
)

/*
	=== Задача №18 ===

	Реализовать структуру-счетчик, которая будет инкрементироваться в конкурентной среде.
	По завершению программа должна выводить итоговое значение счетчика.

*/

type Inc struct {
	sync.RWMutex
	cnt int
}

func main() {
	inc := Inc{cnt: 0}
	var wg sync.WaitGroup
	for i:=0; i<1000; i++ {
		wg.Add(1)
		go func () {
			inc.Lock()
			defer wg.Done()
			defer inc.Unlock()
			inc.cnt++
		}()
	}
	wg.Wait()
	fmt.Println(inc.cnt)
}
