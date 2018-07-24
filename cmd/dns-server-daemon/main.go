package main

import (
	"github.com/astaxie/beego"
	_ "github.com/mojo-zd/dns-server-refresh-plugin/internal/daemon/routers"
	"github.com/mojo-zd/dns-server-refresh-plugin/internal/initialize"
	"github.com/mojo-zd/dns-server-refresh-plugin/internal/variables"
	"github.com/mojo-zd/dns-server-refresh-plugin/pkg"
)

func main() {
	beego.BConfig.CopyRequestBody = true
	if err := pkg.CreateDir(variables.RecordPath); err != nil {
		panic(err.Error())
		return
	}
	initialize.Init()
	beego.Run(variables.DaemonPort)
}
