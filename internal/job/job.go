package job

import (
	"fmt"
	"time"

	"github.com/mojo-zd/dns-server-refresh-plugin/internal/variables"
)

type HealthJob struct {
	URL    string
	worker *Worker
}

func (j HealthJob) HealthCheck() {
	ticker := time.NewTicker(variables.DefaultHealthCheckInterval)
	go func() {
		for {
			select {
			case <-ticker.C:
				fmt.Println("corn task start . . .")
			}
		}
	}()
}
