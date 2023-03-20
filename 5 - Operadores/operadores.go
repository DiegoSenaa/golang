package main

import "fmt"

func main() {
	soma := 1 + 2
	sub := 1 - 2
	div := 10 / 4
	multi := 10 * 5
	mod := 10 % 2

	fmt.Println(soma, sub, div, multi, mod)

	fmt.Println(1 > 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 != 2)
	fmt.Println(1 == 2)
	fmt.Println(1 <= 2)
	fmt.Println(1 < 2)
	fmt.Println(1 <= 2)

	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true || false)

	num := 10
	num++
	fmt.Println(num)
	num += 15
	fmt.Println(num)
	num--
	fmt.Println(num)
	num -= 15
	fmt.Println(num)
	num *= 3
	fmt.Println(num)
	num /= 15
	fmt.Println(num)
	num %= 3
	fmt.Println(num)
}
