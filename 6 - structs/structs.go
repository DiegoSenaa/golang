package main

import "fmt"

type usuario struct {
	nome     string
	idade    int8
	endereco endereco
}

type endereco struct {
	logradouro string
	numero     uint8
}

func main() {

	endereco := endereco{"Rua A", 18}

	usuario1 := usuario{"Davi", 21, endereco}
	fmt.Println(usuario1)

	var usuario2 usuario
	usuario2.nome = "Diego"
	usuario2.idade = 20
	fmt.Println(usuario2)

	usuario3 := usuario{nome: "Diego"}
	fmt.Println(usuario3)
}
