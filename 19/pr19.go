package main

/*
Разработать программу, которая переворачивает подаваемую на ход
строку (например: «главрыба — абырвалг»). Символы могут быть unicode.
*/

import "fmt"

//функция реверса
func reverse(str string) (result string) {
	//идёт запись строки в обратном порядке
	for _, s := range str {
		result = string(s) + result
	}
	return
}

func main() {

	str := "test"
	fmt.Println(str)
	fmt.Println(reverse(str))
}
