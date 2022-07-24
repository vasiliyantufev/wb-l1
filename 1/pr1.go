package main

import "fmt"

/*
task #1
Дана структура Human (с произвольным набором полей и методов).
Реализовать встраивание методов в структуре Action от родительской структуры Human (аналог наследования).
*/

type Human struct {
	name string
	age  int
	sex  bool
}

func (h *Human) Hi() {
	fmt.Println("Hi kitty!")
}

type Action struct {
	Human
}

func main() {
	Action := new(Action)
	Action.Hi()
}
