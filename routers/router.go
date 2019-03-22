package routers

import (
	"github.com/slzm40/gogate/controllers/elinkctls"
	"github.com/slzm40/gogate/controllers/webctls"
	"github.com/slzm40/gomo/elink"
	"github.com/slzm40/gomo/elink/channel/ctrl"

	"github.com/astaxie/beego"
)

// web router setting
func init() {
	beego.Router("/", &webctls.HomeController{})
	beego.Router("/login/:id([0-9]+)", &webctls.LoginController{})
}

// elink router setting
func init() {
	elink.Router(ctrl.ChannelCtrl, "devices.@", &elinkctls.DevicesCtrlController{})
	elink.Router(ctrl.ChannelCtrl, "devices.command.@", &elinkctls.DevCommandCtrlController{})
	elink.Router(ctrl.ChannelCtrl, "zigbee.network", &elinkctls.ZbNetworkCtrlController{})
}
