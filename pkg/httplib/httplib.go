package httplib

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	netUrl "net/url"
	"time"

	"github.com/Sirupsen/logrus"
)

var defaultTimeout = time.Second * 15

type HttpRestTemplate struct {
	method  string
	query   map[string]interface{}
	body    interface{}
	header  map[string]string
	timeout time.Duration
}

func NewHttpRestTemplate() *HttpRestTemplate {
	return &HttpRestTemplate{timeout: defaultTimeout, query: make(map[string]interface{}), header: map[string]string{"ContentType": "application/json"}}
}

// SetTimeout ...
func (h *HttpRestTemplate) SetTimeout(timeout time.Duration) *HttpRestTemplate {
	h.timeout = timeout
	return h
}

// Query query param builder
func (h *HttpRestTemplate) Query(query map[string]interface{}) *HttpRestTemplate {
	h.query = query
	return h
}

// Builder param builder
func (h *HttpRestTemplate) Builder(query map[string]interface{}, body interface{}) *HttpRestTemplate {
	h.query = query
	h.body = body
	return h
}

// Body body builder
func (h *HttpRestTemplate) Body(body interface{}) *HttpRestTemplate {
	h.body = body
	return h
}

// Header header builder
func (h *HttpRestTemplate) Header(header map[string]string) *HttpRestTemplate {
	h.header = header
	return h
}

// Reset 重置 HttpRestTemplate对象
func (h *HttpRestTemplate) Reset() *HttpRestTemplate {
	return &HttpRestTemplate{query: make(map[string]interface{}), header: make(map[string]string)}
}

func (h *HttpRestTemplate) Method(method string) *HttpRestTemplate {
	h.method = method
	return h
}

// Do
func (h *HttpRestTemplate) Do(url string, out interface{}) (err error) {
	var response *http.Response
	var bytes []byte

	switch h.method {
	case http.MethodPost:
		response, err = h.Post(url)
	case http.MethodPut:
		response, err = h.Put(url)
	case http.MethodGet:
		response, err = h.Get(url)
	case http.MethodDelete:
		response, err = h.Delete(url)
	default:
		response, err = h.Get(url)
	}

	if out == nil {
		return
	}

	defer response.Body.Close()
	if bytes, err = ioutil.ReadAll(response.Body); err != nil {
		logrus.Errorf("read from response failed, error is %s", err.Error())
		return
	}

	switch out.(type) {
	case string:
		out = string(bytes)
	default:
		err = json.Unmarshal(bytes, out)
	}
	return
}

// Get
func (h *HttpRestTemplate) Get(url string) (response *http.Response, err error) {
	h.method = http.MethodGet
	var queries string
	for key, value := range h.query {
		queries += fmt.Sprintf("&%s=%s", key, value)
	}

	if len(queries) > 0 {
		queries = queries[1:]
		url = fmt.Sprintf("%s?%s", url, queries)
	}

	response, err = request(h, url)
	return
}

// Post
func (h *HttpRestTemplate) Post(url string) (response *http.Response, err error) {
	h.method = http.MethodPost
	response, err = request(h, url)
	return
}

// Put
func (h *HttpRestTemplate) Put(url string) (*http.Response, error) {
	h.method = http.MethodPut
	return request(h, url)
}

// Delete
func (h *HttpRestTemplate) Delete(url string) (*http.Response, error) {
	h.method = http.MethodDelete
	return request(h, url)
}

func bodyHandle(body interface{}) (reader io.Reader, err error) {
	if body == nil {
		return
	}
	switch body.(type) {
	case string:
		reader = bytes.NewBufferString(body.(string))
	case []byte:
		reader = bytes.NewBuffer(body.([]byte))
	default:
		var bodyBytes []byte
		if bodyBytes, err = json.Marshal(body); err != nil {
			logrus.Errorf("json marshal failed, error is %s", err)
		} else {
			reader = bytes.NewBuffer(bodyBytes)
		}
	}
	return

}

func newRequest(h *HttpRestTemplate, url string, reader io.Reader) (response *http.Response, err error) {
	u, err := netUrl.Parse(url)
	if err != nil {
		logrus.Errorf("parse url failed,err:%s", err.Error())
		return
	}

	req := &http.Request{
		Method:     h.method,
		URL:        u,
		Header:     make(http.Header),
		Proto:      "HTTP/1.1",
		ProtoMajor: 1,
		ProtoMinor: 1,
	}

	if reader != nil {
		req.Body = ioutil.NopCloser(reader)
	}

	if len(h.header) > 0 {
		for k, v := range h.header {
			req.Header.Add(k, v)
		}
	}

	client := http.Client{Timeout: h.timeout}
	response, err = client.Do(req)
	return
}

func request(h *HttpRestTemplate, url string) (response *http.Response, err error) {
	reader, err := bodyHandle(h.body)
	if err != nil {
		return
	}
	if response, err = newRequest(h, url, reader); err != nil {
		logrus.Errorf("request %s failed, error is %s", url, err.Error())
	}
	return
}
