package routers

import (
	"github.com/thinkgos/gogate/apps/elinkch/ctrl"
	"github.com/thinkgos/gogate/controllers/elinkctls"
	"github.com/thinkgos/gogate/pkg/elink"
)

// elink router setting
func init() {
	elink.Router(ctrl.ChannelCtrl, "system.user", &elinkctls.SysUserController{})
	elink.Router(ctrl.ChannelCtrl, "gateway.upgrade", &elinkctls.GatewayUpgradeController{})
	elink.Router(ctrl.ChannelCtrl, "gateway.infos", &elinkctls.GatewayInfosController{})
	elink.Router(ctrl.ChannelCtrl, "devices.@", &elinkctls.DevicesController{})
	elink.Router(ctrl.ChannelCtrl, "device.commands.@", &elinkctls.DevCommandController{})
	elink.Router(ctrl.ChannelCtrl, "device.propertys.@", &elinkctls.DevPropertysController{})
	elink.Router(ctrl.ChannelCtrl, "zigbee.network", &elinkctls.ZbNetworkController{})
}
