package main

/*

15 К каким негативным последствиям может привести данный фрагмент кода, и как это исправить?
Приведите корректный пример реализации.

var justString string
func someFunc() {
    v := createHugeString(1 << 10)
    justString = v[:100]
}

func main() {
so  meFunc()
}
*/

import "fmt"

var justString string

func main() {
	someFunc()
}

func someFunc() {
	v := "zYj@$AdFx7UzMspXRV6DA0aIx*GbT0velZCsS?s5pPb%4oD#GoT7reKx*IU@%#q0rZlWN?Z}dWvQ6y$yvLyw|A1QoyHXB3ApHb7"
	r := []rune(v)
	justString = string(r[:20])
	fmt.Println(justString)
}