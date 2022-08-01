package main

import "fmt"

/*
Разработать программу, которая в рантайме способна определить тип
переменной: int, string, bool, channel из переменной типа interface{}.

*/

func main() {

	var t interface{}
	t = 1
	typeDefine(t)
}

// определяем тип(int, string, bool, channel) переданной переменной.
func typeDefine(i interface{}) {

	switch i.(type) {
	case int:
		fmt.Println("тип int")
	case string:
		fmt.Println("тип string")
	case bool:
		fmt.Println("тип bool")
	case chan int:
		fmt.Println("тип chan int")
	case chan string:
		fmt.Println("тип chan string")
	case chan bool:
		fmt.Println("тип chan bool")
	default:
		fmt.Println("Не распознан")
	}
}
