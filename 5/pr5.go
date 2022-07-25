package main

import (
	"fmt"
	"time"
)

/*
Разработать программу, которая будет последовательно отправлять значения
в канал, а с другой стороны канала — читать.
По истечению N секунд программа должна завершаться.
*/

func main() {

	const N = 10
	ch := make(chan int)
	go print(ch)
	go send(ch)
	//go print(ch)
	time.Sleep(time.Duration(N * time.Second))
	close(ch)
}

func print(ch <-chan int) {
	for v := range ch {
		fmt.Println(v)
	}
}

func send(ch chan<- int) {
	i := 0
	for {
		ch <- i
		i++
	}

}
