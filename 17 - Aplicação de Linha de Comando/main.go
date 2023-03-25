package main

import (
	"fmt"
	"linha-de-comando/app"
	"log"
	"os"
)

func main() {
	fmt.Println("Iniciando aplicação")

	aplicacao := app.Gerar()
	if er := aplicacao.Run(os.Args); er != nil {
		log.Fatal(er)
	}
}
