package main

import "fmt"

// Реализовать все возможные способы остановки выполнения горутины.

func main() {
	quit := make(chan bool)

	go func() {
		fmt.Println("start")
		for {
			select {
			case <-quit:
				return
			default:

			}
		}
	}()
	//..
	fmt.Println("Enter")
	quit <- true
	fmt.Println("Exit")

}
