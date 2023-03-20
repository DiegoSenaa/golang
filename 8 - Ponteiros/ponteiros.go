package main

import "fmt"

func main() {
	var var1 int = 10
	var var2 int = var1

	fmt.Println(var1, var2)

	var1++
	fmt.Println(var1, var2)

	//Ponteiro é ref de memória
	var var3 int = 100
	var ponteiro *int
	var3 = 100
	ponteiro = &var3

	fmt.Println(var3, ponteiro)
	//Desreferênciado o ponteiro
	fmt.Println(var3, *ponteiro)

	var3 = 150
	fmt.Println(var3, ponteiro)
	//Desreferênciado o ponteiro
	fmt.Println(var3, *ponteiro)
}
