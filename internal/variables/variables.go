package variables

import (
	"fmt"
	"time"
)

type Action string

const DaemonPort = ":10106"

var (
	Add    Action = "add"
	Delete Action = "delete"
	Clear  Action = "clear"

	ConfFile                   = "records.json"
	DefaultHealthCheckInterval = time.Duration(time.Second * 5)

	RecordPath                 = "/tmp/dns-server-refresh"
	DefaultDnsRefreshDaemonURL = fmt.Sprintf("%s%s", "http://127.0.0.1", DaemonPort)
)

// environment variables
var (
	RECORD_PATH            = "RECORD_PATH"
	DNS_REFRESH_DAEMON_URL = "DNS_REFRESH_DAEMON_URL"
)
