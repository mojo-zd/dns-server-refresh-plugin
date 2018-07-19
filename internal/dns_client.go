package internal

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/Sirupsen/logrus"
	"github.com/mojo-zd/dns-server-refresh-plugin/internal/variables"
	"github.com/mojo-zd/dns-server-refresh-plugin/pkg"
)

var (
	HttpPrefix  = "http://"
	HttpsPrefix = "https://"
)

type DnsClient struct {
	client
}

func NewDnsClient(confPath string) *DnsClient {
	return &DnsClient{client{confPath: confPath}}
}

type client struct {
	confPath string // conf path
}

// Clear ...
func (c *DnsClient) Clear() (err error) {
	err = pkg.DeleteFile(c.conf())
	pkg.ClearFile(c.conf())
	return
}

// Write ...
func (c *DnsClient) Write(url string, action variables.Action) (err error) {
	m, err := c.readFile()
	if err != nil {
		logrus.Errorf("write record failed, err: %s", err.Error())
		return
	}

	var key string
	if ss := c.decompose(url); len(c.decompose(url)) > 0 {
		key = ss[0]
	} else {
		logrus.Warningf("decompose url %s failed, check url correctness", url)
		return
	}
	switch action {
	case variables.Add:
		m[key] = url
	case variables.Delete:
		delete(m, key)
	}

	err = pkg.WriteFile(c.conf(), m)
	return
}

func (c client) decompose(url string) (ss []string) {
	s := url
	if strings.HasPrefix(url, HttpPrefix) {
		s = url[strings.Index(url, HttpPrefix)+len(HttpPrefix):]
	}

	if strings.HasPrefix(url, HttpsPrefix) {
		s = url[strings.Index(url, HttpsPrefix)+len(HttpsPrefix):]
	}
	ss = strings.SplitN(s, "/", 2)
	return
}

func (c client) conf() (file string) {
	if len(c.confPath) == 0 {
		file = variables.ConfFile
	} else {
		if strings.HasSuffix(c.confPath, "/") {
			file = fmt.Sprintf("%s%s", c.confPath, variables.ConfFile)
		} else {
			file = fmt.Sprintf("%s/%s", c.confPath, variables.ConfFile)
		}
	}
	return
}

func (c client) readFile() (m map[string]interface{}, err error) {
	m = map[string]interface{}{}
	bytes, err := pkg.ReadFile(c.conf())
	if err != nil {
		return
	}

	if len(bytes) == 0 {
		return
	}

	err = json.Unmarshal(bytes, &m)
	return
}
