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

}
