package job

import (
	"fmt"
	"testing"
	"time"

	"github.com/mojo-zd/dns-server-refresh-plugin/internal/job"
)

func TestWorker(t *testing.T) {
	worker := job.NewWorker()
	worker.Run(map[string]interface{}{"118.31.50.65": "http://118.31.50.55/i18n/lang/en-us-lang.json", "127.0.0.1": "http://127.0.0.1"})
	timer := time.NewTicker(time.Second * 21)
	worker.Watch(func(args ...interface{}) {
		fmt.Println("失败调用", args)
	}, "dns server record filed location", "dns process name")

	for {
		select {
		case <-timer.C:
		}
	}
}
