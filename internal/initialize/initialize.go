package initialize

import (
	"github.com/mojo-zd/dns-server-refresh-plugin/internal"
	"github.com/mojo-zd/dns-server-refresh-plugin/internal/job"
	"github.com/mojo-zd/dns-server-refresh-plugin/pkg"
)

var (
	DnsClient = internal.NewDnsClient(pkg.RecordPath())
	Worker    = job.NewWorker()
)

func Init() {
	records, err := DnsClient.Read()
	if err != nil {
		return
	}

	Worker.Run(records)
}
