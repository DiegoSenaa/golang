package main

import (
	"fmt"
	"reflect"
)

func main() {
	var array1 [5]int
	array1[0] = 1
	fmt.Println(array1)

	array2 := [5]int8{1, 2, 3, 4, 5}
	fmt.Println(array2)

	array3 := [...]int{1, 2, 3, 4, 5, 6}
	fmt.Println(array3)

	slice1 := []int{1, 2, 3, 4}
	fmt.Println(slice1)
	slice1 = append(slice1, 100)
	fmt.Println(slice1)

	fmt.Println(reflect.TypeOf(slice1))
	fmt.Println(reflect.TypeOf(array1))

	slice2 := array2[1:3]
	fmt.Println(slice2)
	array2[1] = 100
	fmt.Println(slice2)

	//Arrays internos

	slice3 := make([]float32, 10, 11)
	fmt.Println(slice3)
	fmt.Println(cap(slice3)) //capacidade
	fmt.Println(len(slice3)) // tamanho

	slice3 = append(slice3, 5)
	slice3 = append(slice3, 6)

	fmt.Println(cap(slice3)) //capacidade
	fmt.Println(len(slice3)) // tamanho

	slice4 := make([]float32, 5)
	fmt.Println(slice4)
	fmt.Println(cap(slice4)) //capacidade
	fmt.Println(len(slice4)) // tamanho

	slice3 = append(slice4, 10)

	fmt.Println(cap(slice4)) //capacidade
	fmt.Println(len(slice4)) // tamanho
}
