package variables

import "time"

type Action string

var (
	Add    Action = "add"
	Delete Action = "delete"
	Clear  Action = "clear"

	ConfFile                   = "conf.json"
	DefaultHealthCheckInterval = time.Duration(time.Second * 5)
)
