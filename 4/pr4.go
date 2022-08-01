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

	// Создаем канал для получения данных, отправляемых воркерам.
	input := make(chan int)

	// Запускаем n воркеров через горутины, передавая номер воркера i и канал ch.
	for i := 0; i < n; i++ {
		go worker(i, input)
	}

	for {
		input <- time.Now().Second()
		time.Sleep(time.Second)
	}
}

// Принимает воркер и канал, данные читаются из канала и записываются в number
func worker(worker int, in <-chan int) {
	for {
		number := <-in
		fmt.Printf("goroutine #%d: value: #%d\n", worker, number)
	}
}
