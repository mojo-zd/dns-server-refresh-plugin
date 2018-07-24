package controllers

import (
	"github.com/mojo-zd/dns-server-refresh-plugin/internal/initialize"
	"github.com/mojo-zd/dns-server-refresh-plugin/internal/variables"
)

// Daemon API
type DnsDaemonController struct {
	ConcreteController
}

type Record struct {
	Name string
}

//@Title Add
//@Description add record
//@Param   key	body Record true  "record"
//@Success 200 {object} Record
//@Failure 400 bad params
//@router / [post]
func (d *DnsDaemonController) Add() {
	record := &Record{}
	d.ErrorHandler(d.Unmarshal(record))

	err := initialize.DnsClient.Write(record.Name, variables.Add)
	d.ErrorHandler(err)

	initialize.Worker.AddJob(record.Name)
	d.ServeJSON()
}

//@Title Delete
//@Description Delete record
//@Param  url  query string true  "record key"
//@Success 200 int http.code
//@Failure 400 bad params
//@router / [delete]
func (d *DnsDaemonController) Delete() {
	url := d.GetString("url")
	if len(url) == 0 {
		return
	}

	initialize.Worker.Remove(url)
	d.ServeJSON()
}

//@Title Clear
//@Description clear record
//@Success 200 {object} Record
//@Failure 400 bad params
//@router /clear [post]
func (d *DnsDaemonController) Clear() {
	initialize.DnsClient.Clear()
	initialize.Worker.Stop()
	d.ServeJSON()
}
