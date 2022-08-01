package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	for i := 0; i < 5; i++ {
		wg.Add(1)
		// нужно принимать указатель на WaitGroup
		// go func(wg * sync.WaitGroup, i int) {
		go func(wg *sync.WaitGroup, i int) {
			fmt.Println(i)
			wg.Done()
			// передается wg как значение
			//}(wg, i)
		}(&wg, i)
	}
	wg.Wait()
	fmt.Println("exit")
}
