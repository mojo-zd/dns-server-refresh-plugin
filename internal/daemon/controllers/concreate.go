package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/astaxie/beego"
)

type ConcreteController struct {
	beego.Controller
}
type ResponseBody struct {
	Code    int
	Message string
}

// Unmarshal decode data from body
func (c *ConcreteController) Unmarshal(i interface{}) (err error) {
	err = json.Unmarshal(c.Ctx.Input.RequestBody, &i)
	return
}

// ErrorHandler ...
func (c *ConcreteController) ErrorHandler(err error, args ...interface{}) {
	if err == nil {
		return
	}
	body := &ResponseBody{}
	switch len(args) {
	case 1:
		body.Code = args[0].(int)
	case 2:
		body.Code = args[0].(int)
		body.Message = args[1].(string)
	default:
		body.Code = http.StatusBadRequest
		body.Message = fmt.Sprintf("参数错误 err: %s", err.Error())
	}

	c.write(body)
	panic(err.Error())
}

func (c *ConcreteController) write(body *ResponseBody) {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "application/json; charset=utf-8")
	c.Ctx.ResponseWriter.WriteHeader(body.Code)
	b, _ := json.Marshal(body)
	c.Ctx.ResponseWriter.Write(b)
}
