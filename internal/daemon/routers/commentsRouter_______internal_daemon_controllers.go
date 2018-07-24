package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["github.com/mojo-zd/dns-server-refresh-plugin/internal/daemon/controllers:DnsDaemonController"] = append(beego.GlobalControllerRouter["github.com/mojo-zd/dns-server-refresh-plugin/internal/daemon/controllers:DnsDaemonController"],
		beego.ControllerComments{
			Method:           "Add",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["github.com/mojo-zd/dns-server-refresh-plugin/internal/daemon/controllers:DnsDaemonController"] = append(beego.GlobalControllerRouter["github.com/mojo-zd/dns-server-refresh-plugin/internal/daemon/controllers:DnsDaemonController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           `/`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["github.com/mojo-zd/dns-server-refresh-plugin/internal/daemon/controllers:DnsDaemonController"] = append(beego.GlobalControllerRouter["github.com/mojo-zd/dns-server-refresh-plugin/internal/daemon/controllers:DnsDaemonController"],
		beego.ControllerComments{
			Method:           "Clear",
			Router:           `/clear`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Params:           nil})

}
