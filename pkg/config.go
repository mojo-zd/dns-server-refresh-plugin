package pkg

import (
	"os"

	"github.com/mojo-zd/dns-server-refresh-plugin/internal/variables"
)

func RecordPath() string {
	return get(variables.RECORD_PATH, variables.RecordPath)
}

func DnsDaemonURL() string {
	return get(variables.DNS_REFRESH_DAEMON_URL, variables.DefaultDnsRefreshDaemonURL)
}

func get(name, def string) string {
	if v := os.Getenv(name); len(v) != 0 {
		return v
	}
	return def
}
