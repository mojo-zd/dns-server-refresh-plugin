package internal

import "github.com/mojo-zd/dns-server-refresh-plugin/internal/variables"

type DnsCommand struct {
	Action variables.Action // address operation eg: create 、delete 、clear
	URL    *string          // health check url eg: http://xx.xx.xxx.xx/api/health
}

func DnsHandler(cmd DnsCommand, client *DnsClient) (err error) {
	switch cmd.Action {
	case variables.Add:
		if cmd.URL == nil {
			return
		}
		err = client.Write(*cmd.URL, cmd.Action)
	case variables.Delete:
		if cmd.URL == nil {
			return
		}
		err = client.Write(*cmd.URL, cmd.Action)
	case variables.Clear:
		err = client.Clear()
	}
	return
}
