package internal

import (
	"fmt"
	"testing"
	"time"

	"github.com/mojo-zd/dns-server-refresh-plugin/internal"
	"github.com/mojo-zd/dns-server-refresh-plugin/internal/variables"
)

func TestDnsClient(t *testing.T) {
	client := internal.NewDnsClient("")
	client.Write("http://10.0.0.151:80/api/health/check", variables.Add)
	//client.Clear()
}

func TestTimer(t *testing.T) {
	ticker := time.NewTicker(time.Second * 5)
	i := 0
	for {
		select {
		case <-ticker.C:
			i += 5
			fmt.Println("print ", i)
		}
	}
}
