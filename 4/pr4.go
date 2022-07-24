package main

import (
	"fmt"
	"time"
)

/*
Реализовать постоянную запись данных в канал (главный поток).
Реализовать набор из N воркеров, которые читают произвольные данные из канала и выводят в stdout.
Необходима возможность выбора количества воркеров при старте.

Программа должна завершаться по нажатию Ctrl+C.
Выбрать и обосновать способ завершения работы всех воркеров.
*/

func main() {
	var n int
	fmt.Println("Введите кол-во горутин:")
	fmt.Scan(&n)

	input := make(chan int)

	for i := 0; i < n; i++ {
		go worker(i, input)
	}

	for {
		input <- time.Now().Second()
		time.Sleep(time.Second)
	}
}

func worker(worker int, in <-chan int) {
	for {
		number := <-in
		fmt.Printf("goroutine #%d: value: #%d\n", worker, number)
	}
}
