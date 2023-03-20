package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	idade     int8
	altura    uint8
}

type estudante struct {
	pessoa
	curso     string
	faculdade string
}

func main() {
	fmt.Println("Herança")

	p1 := pessoa{"Joao", "Pedro", 20, 160}
	fmt.Println(p1)

	e1 := estudante{p1, "Eng", "USP"}
	fmt.Println(e1.nome)
}
