package main

import "fmt"

func somar(n1 int8, n2 int8) int8 {
	return n1 + n2
}

func main() {
	soma := somar(10, 20)
	fmt.Println(soma)

	f := func(txt string) string {
		fmt.Println(txt)
		return txt
	}

	f("print")
	retorno := f("retorno")
	fmt.Println(retorno)

	soma, sub := calculosMatematicos(1, 1)
	fmt.Println(soma, sub)

	soma2, _ := calculosMatematicos(1, 1)
	fmt.Println(soma2)
}

func calculosMatematicos(n1, n2 int8) (int8, int8) {
	soma := n1 + n2
	sub := n1 - n2
	return soma, sub
}
