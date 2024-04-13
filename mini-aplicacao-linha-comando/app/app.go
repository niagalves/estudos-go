package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

func buscarIps(c *cli.Context) {
	host := c.String("host")

	// pacote net
	// vai buscar os ips
	ips, err := net.LookupIP(host)

	if err != nil {
		log.Fatal(err)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}
}

func buscaServidores(c *cli.Context) {
	host := c.String("host")

	servidores, err := net.LookupNS(host)

	if err != nil {
		log.Fatal(err)
	}

	for _, servidor := range servidores {
		fmt.Println(servidor.Host)
	}

}

// Gerar retorna a aplicação de linha de comando
// pronta para ser executada
// pra executar o comando digite
// go run main.go ip --host seu site
// go run main.go servidores --host seu site
func Gerar() *cli.App {
	app := cli.NewApp()
	app.Name = "Aplicação de linha de comando"
	app.Usage = "Busca IPs e Nomes de Servidor na internet"

	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "github.com",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:   "ip",
			Usage:  "Busca ips de endereços na internet",
			Flags:  flags,
			Action: buscarIps,
		},
		{
			Name:   "servidores",
			Usage:  "Busca o nome dos servidores",
			Flags:  flags,
			Action: buscaServidores,
		},
	}
	return app
}
