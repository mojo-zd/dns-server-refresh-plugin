package main

import (
	"os"

	"github.com/mojo-zd/dns-server-refresh-plugin/internal"
	"github.com/mojo-zd/dns-server-refresh-plugin/internal/variables"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	app = kingpin.New("dns-refresh", "A command-line to check domain health and refresh dns server.")
)

func main() {
	dnsClient := internal.NewDnsClient("")
	cmd := internal.DnsCommand{}

	addCommand := app.Command("add", "add health check task")
	addURL := addCommand.Flag("url", "the url of health check").String()

	deleteCommand := app.Command("delete", "delete health check task")
	deleteURL := deleteCommand.Flag("url", "the url of health check").String()

	cmd.Action = variables.Action(kingpin.MustParse(app.Parse(os.Args[1:])))

	switch cmd.Action {
	case variables.Action(addCommand.FullCommand()):
		cmd.URL = addURL
	case variables.Action(deleteCommand.FullCommand()):
		cmd.URL = deleteURL
	}

	internal.DnsHandler(cmd, dnsClient)
}
