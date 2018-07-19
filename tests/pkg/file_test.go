package pkg

import (
	"fmt"
	"strings"
	"testing"

	"github.com/mojo-zd/dns-server-refresh-plugin/pkg"
)

func TestReadFile(t *testing.T) {
	f, err := pkg.ReadFile("mojo.json")
	if err != nil {
		t.Error(err.Error())
		return
	}
	t.Log(string(f))
}

func TestDecompose(t *testing.T) {
	a := "127.0.0.1"
	fmt.Println(decompose(a))
}

func decompose(url string) (s string) {
	s = url
	if strings.HasPrefix(url, "http://") {
		s = url[strings.Index(url, "http://")+len("http://"):]
	}

	if strings.HasPrefix(url, "https://") {
		s = url[strings.Index(url, "https://")+len("https://"):]
	}

	ss := strings.SplitN(s, "/", 2)
	fmt.Println(ss)
	return
}
