package main

import "fmt"

/*
Поменять местами два числа без создания временной переменной.
*/

func main() {

	x := 0
	y := 1

	//меняем местами х и у
	x, y = y, x

	fmt.Println("x=%d, y=%d \n", x, y)

}
