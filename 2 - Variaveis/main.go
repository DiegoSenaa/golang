package main

import "fmt"

func main() {

	//Declaração variável explicita
	var var1 string = "Var1"
	fmt.Println(var1)

	//Declaração variável por inferência.
	var2 := "Var2"
	fmt.Println(var2)

	//Declaração multipla de variáveis explicita
	var (
		var3 string = "blah"
		var4 string = "blah"
	)
	fmt.Println(var3, var4)

	//Declaração multipla de variáveis por inferência.
	var5, var6 := "bar", "bar"
	fmt.Println(var5, var6)

	//Declaração de constantes explicita
	const const1 string = "Foo"
	fmt.Println(const1)

	//Declaração constate por inferência.

}
