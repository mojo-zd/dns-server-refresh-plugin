package internal

import (
	"fmt"

	"github.com/mojo-zd/dns-server-refresh-plugin/internal/job"
	"github.com/mojo-zd/dns-server-refresh-plugin/internal/variables"
	"github.com/mojo-zd/dns-server-refresh-plugin/pkg"
	"github.com/mojo-zd/dns-server-refresh-plugin/pkg/httplib"
)

var (
	worker                *job.Worker
	daemonRecordApiPrefix = "/api/records"
)

func init() {
	worker = job.NewWorker()
}

type DnsCommand struct {
	Action variables.Action // address operation eg: create 、delete 、clear
	URL    *string          // health check url eg: http://xx.xx.xxx.xx/api/health
}

// DnsHandler ...
func DnsHandler(cmd DnsCommand, client *DnsClient) (err error) {
	switch cmd.Action {
	case variables.Add:
		if cmd.URL == nil {
			return
		}
		httplib.NewHttpRestTemplate().Body(map[string]interface{}{"name": *cmd.URL}).Post(fmt.Sprintf("%s%s", pkg.DnsDaemonURL(), daemonRecordApiPrefix))
	case variables.Delete:
		if cmd.URL == nil {
			return
		}
		httplib.NewHttpRestTemplate().Query(map[string]interface{}{"url": *cmd.URL}).Delete(fmt.Sprintf("%s%s", pkg.DnsDaemonURL(), daemonRecordApiPrefix))
	case variables.Clear:

	}
	return
}
