package main

import "fmt"

/*
task #1
Дана структура Human (с произвольным набором полей и методов).
Реализовать встраивание методов в структуре Action от родительской структуры Human (аналог наследования).
*/

// Родительская структура
type Human struct {
	name string
	age  int
	sex  bool
}

//метод объявлен для типа Human
func (h *Human) Hi() {
	fmt.Println("Hi kitty!")
}

//
type Action struct { //встраиваем Human в Action дочернюю структуру
	Human
}

func main() {
	// Создание экземпляра.
	Action := new(Action)
	// дочерняя структура может напрямую обращаться к переменным базовой структуры.
	Action.Hi()
}
