package main

import "fmt"

//для того чтобы изменить значение переменной (a) указателя p,
//нужно его разыменовать *p, а потом присвоить значение.
//тогда a присвоется новое значение

func update(p *int) {
	b := 2
	//p = &b
	*p = b
}

func main() {
	var (
		a = 1
		p = &a
	)
	fmt.Println(*p)
	update(p)
	fmt.Println(*p)
}
