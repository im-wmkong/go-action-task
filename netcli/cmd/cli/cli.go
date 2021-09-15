package main

import (
	"fmt"
	"github.com/urfave/cli"
	"log"
	"net"
	"os"
)

func ns(context *cli.Context) error {
	nss, err := net.LookupNS(context.Args().First())
	if err != nil {
		return err
	}
	for _, ns := range nss {
		fmt.Println(*ns)
	}
	return nil
}

func cname(context *cli.Context) error {
	cname, err := net.LookupCNAME(context.Args().First())
	if err != nil {
		return err
	}
	fmt.Println(cname)
	return nil
}

func ip(context *cli.Context) error {
	ips, err := net.LookupIP(context.Args().First())
	if err != nil {
		return err
	}
	for _, ip := range ips {
		fmt.Println(ip)
	}
	return nil
}

func help(context *cli.Context) error {
	cmd := getCommands()
	for _, v := range cmd {
		fmt.Println(v.Name)
	}
	return nil
}

func getCommands() cli.Commands {
	return cli.Commands{
		{
			Name:   "ns",
			Usage:  "Looks Up the NameServers for a Particular Host",
			Action: ns,
		},
		{
			Name:   "cname",
			Usage:  "Looks up the CNAME for a particular host",
			Action: cname,
		},
		{
			Name:   "ip",
			Usage:  "Looks up the IP addresses for a particular host",
			Action: ip,
		},
		{
			Name:    "help",
			Aliases: []string{"h"},
			Usage:   "Shows a list of commands or help for one command",
			Action:  help,
		},
	}
}

func main() {
	app := cli.NewApp()
	app.Name = "Website Lookup CLI"
	app.Usage = "Let's you query IPs, CNAMEs and Name Servers!"
	app.Version = "0.0.0"
	app.Commands = getCommands()

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
