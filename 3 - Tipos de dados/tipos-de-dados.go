package main

import (
	"errors"
	"fmt"
)

func main() {

	// vai até o int64, se usar inferencia toma como base a arquitetura do pc o mesmo para apenas int, suporta negativos
	var num int16 = 100
	fmt.Println(num)

	//Não suporta numeros negativos, segue a mesma ideia do int
	var num2 uint32 = 100
	fmt.Println(num2)

	//tipos tem alias
	//INT32 = RUNE
	var num3 rune = 1234
	fmt.Println(num3)

	//BYTE = UINT8
	var num4 byte = 123
	fmt.Println(num4)

	//float32 e float64, aceitam ponto flutuante, se não indicar um dos tipos ele tbm pega pela arquitetura.
	var numReal float32 = 123.45
	fmt.Println(numReal)

	var numReal2 float64 = 12300000.45
	fmt.Println(numReal2)

	//String é uma cadeia de caracteris, não existe char de fato
	var str string = "Texto"
	fmt.Println(str)

	str2 := "Texto2"
	fmt.Println(str2)

	//Se usar aspas simples ele devolve o code da tabela ascii
	char := 'B'
	fmt.Println(char)

	//Valor Zero, são valores dados para as variaveis, caso não seja atribuido um valor inicial

	var valorZeroString string
	fmt.Println(valorZeroString)

	var valorZeroInt int
	fmt.Println(valorZeroInt)

	//booleano, só aceita true ou false
	var bool1 = true
	fmt.Println(bool1)

	//erro
	var erro error = errors.New("Erro interno")
	fmt.Println(erro)

}
