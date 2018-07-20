package job

import (
	"net/http"
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/mojo-zd/dns-server-refresh-plugin/internal/variables"
	"github.com/mojo-zd/dns-server-refresh-plugin/pkg/httplib"
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
				response, err := httplib.NewHttpRestTemplate().Get(j.URL)
				if err != nil || response.StatusCode < http.StatusOK || response.StatusCode >= http.StatusBadRequest {
					logrus.Warnf("%s not health url", j.URL)
					j.worker.stop <- j.URL
					return
				}
			}
		}
	}()
}
