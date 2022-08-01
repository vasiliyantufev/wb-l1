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
	someFunc()
}
*/

import "fmt"

func main() {
	//v := "zYj@$AdFx7UzMspXRV6DA0aIx*GbT0velZCsS?s5pPb%4oD#GoT7reKx*IU@%#q0rZlWN?Z}dWvQ6y$yvLyw|A1QoyHXB3ApHb7"
	v1 := "ЯМПВРРОЩ_ырдщ00使 и 便 · 使– shǐ使 и 便 · 使– shǐ使 и 便 · 使– shǐ使 и 便 · 使– shǐ使 и 便 · 使– shǐ"
	justString := someFunc(v1)
	fmt.Println(justString)
}

//go кодирует строку как набор байт, поэтому если используюьтся специфические значения, то нужно использовать руны
func someFunc(v string) string {
	r := []rune(v)
	return string(r[:77])
	//return string(v[:77])
}
