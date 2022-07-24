package main

import "fmt"

/*
task #3
Дана последовательность чисел: 2,4,6,8,10. Найти сумму их
квадратов(22+32+42….) с использованием конкурентных вычислений.
*/

func main() {
	array := []int{2, 4, 6, 8, 10}

	result := 0

	for _, a := range array {
		result += a * a
	}
	fmt.Printf("Сумма квадратов = %v\n", result)
}
