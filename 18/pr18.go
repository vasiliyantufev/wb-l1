package main

/*
Реализовать структуру-счетчик, которая будет инкрементироваться в конкурентной среде.
По завершению программа должна выводить итоговое значение счетчика.
*/

import (
	"fmt"
	"sync"
)

type counter struct {
	mx sync.Mutex
	wg sync.WaitGroup
	i  int
}

//создаём структуру для счётчика
func counterStruct() counter {
	return counter{}
}

//увеличивае значение счётчика
func increment(count *counter) {
	//мьютекс блокировка с общим доступом
	count.mx.Lock()
	count.i++
	count.mx.Unlock()
	count.wg.Done()
}

func main() {
	counter := counterStruct()

	var n int
	fmt.Print("Введите число: ")
	fmt.Scan(&n)

	fmt.Println("Значение счетчика до вычислений", counter.i)
	//Под каждую i создаётся горутина
	for i := 0; i < n; i++ {
		counter.wg.Add(1)
		go increment(&counter)
	}
	counter.wg.Wait()
	fmt.Println("Значение счетчика после вычислений", counter.i)
}
