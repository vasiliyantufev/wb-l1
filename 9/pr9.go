package main

import (
	"fmt"
)

/*
Разработать конвейер чисел. Даны два канала: в первый пишутся числа (x) из массива,
во второй — результат операции x*2, после чего данные из второго канала должны выводиться в stdout.
*/

func main() {

	ch1 := make(chan int)
	ch2 := make(chan int)

	array := []int{1, 2, 4, 6, 8}

	go func() {
		for _, val := range array {
			ch1 <- val
		}
		close(ch1)
	}()

	go func() {
		for val := range ch1 {
			ch2 <- val * 2

		}
		close(ch2)
	}()

	for x := range ch2 {
		fmt.Println(x)
	}
}
