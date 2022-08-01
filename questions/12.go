package main

import "fmt"

func main() {
	n := 0
	if true {
		n := 1 // новая локальная переменная
		n++
	}
	fmt.Println(n)
}
