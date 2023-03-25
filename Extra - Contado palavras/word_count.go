package main

import (
	"fmt"
	"strings"
)

func main() {

	s := "Não acho que quem ganhar ou quem perder, nem quem ganhar nem perder, vai ganhar ou perder. Vai todo mundo perder"
	t := strings.ReplaceAll(s, ",", "")
	z := strings.ReplaceAll(t, ".", "")
	words := strings.Fields(z)
	m := make(map[string]int)
	for _, word := range words {
		m[word] += 1
	}
	fmt.Println(m)
}
