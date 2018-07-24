package routers

import (
	"github.com/astaxie/beego"
	"github.com/mojo-zd/dns-server-refresh-plugin/internal/daemon/controllers"
)

func init() {
	ns := beego.NewNamespace("/api",
		beego.NSNamespace("/records",
			beego.NSInclude(
				&controllers.DnsDaemonController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
