package main

import "fmt"

// Реализовать все возможные способы остановки выполнения горутины.

func main() {

	// передаем false (0)
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
	// закрытие канала
	close(quit)

}
