package main

import "fmt"

/*
Дана переменная int64. Разработать программу
которая устанавливает i-й бит в 1 или 0.
*/

func main() {

	var num int64
	fmt.Print("Введите число: ")
	fmt.Scan(&num)
	fmt.Printf("Введенно число: %d в двоичном представлении %b\n", num, num)

	b := false
	i := 2

	if b {
		num |= 1 << i
	} else {
		num &^= 1 << i
	}
	fmt.Printf("Введенно число: %d в двоичном представлении после изменений %b", num, num)
}
