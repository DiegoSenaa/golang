package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

// Gerar vai retornar a apliação cli para ser executada
func Gerar() *cli.App {
	app := cli.NewApp()
	app.Name = "Aplicação de linha de comando."
	app.Usage = "Busca ips e nomes de servidores na internet"

	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "devbook.com.br",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:   "ip",
			Usage:  "busca ips de endereços na internet",
			Flags:  flags,
			Action: buscarIps,
		},
		{
			Name:   "servidores",
			Usage:  "busca nome dos servidores na internet",
			Flags:  flags,
			Action: buscaServidores,
		},
	}

	return app
}

func buscarIps(c *cli.Context) {
	host := c.String("host")

	ips, er := net.LookupIP(host)

	if er != nil {
		log.Fatal(er)
	}

	for index, ip := range ips {
		fmt.Println(index, "-", ip)
	}
}

func buscaServidores(c *cli.Context) {
	host := c.String("host")

	ns, er := net.LookupNS(host)

	if er != nil {
		log.Fatal(er)
	}

	for index, ns := range ns {
		fmt.Println(index, "-", ns)
	}
}
