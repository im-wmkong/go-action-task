package main

import (
	"fmt"
	"github.com/urfave/cli"
	"log"
	"net"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "Website Lookup CLI"
	app.Usage = "Let's you query IPs, CNAMEs and Name Servers!"
	app.Version = "0.0.0"

	app.Commands = cli.Commands{
		{
			Name:  "ns",
			Usage: "Looks Up the NameServers for a Particular Host",
			Action: func(context *cli.Context) error {
				ns, err := net.LookupNS(context.Args().Get(0))
				if err != nil {
					return err
				}
				fmt.Println(ns)
				return nil
			},
		},
		{
			Name:  "cname",
			Usage: "Looks up the CNAME for a particular host",
			Action: func(context *cli.Context) error {
				return nil
			},
		},
		{
			Name:  "ip",
			Usage: "Looks up the IP addresses for a particular host",
			Action: func(context *cli.Context) error {
				return nil
			},
		},
		{
			Name:  "help",
			Usage: "Shows a list of commands or help for one command",
			Action: func(context *cli.Context) error {
				return nil
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
