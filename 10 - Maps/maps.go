package main

import "fmt"

func main() {
	fmt.Println("Maps")

	user := map[string]string{
		"nome":      "Pedro",
		"sobrenome": "Silva",
	}

	fmt.Println(user)

	fmt.Println(user["nome"])

	user2 := map[string]map[string]string{
		"nome": {
			"nome":      "Pedro",
			"sobrenome": "Silva",
		},
		"curso": {
			"nome":   "S.I",
			"campus": "Campus 1",
		},
	}

	fmt.Println(user2)

	delete(user2, "nome")
	fmt.Println(user2)

	user2["signo"] = map[string]string{
		"nome": "Gêmios",
	}
}
