package elinkctls

import (
	"github.com/thinkgos/gogate/apps/elinkch/ctrl"
	"github.com/thinkgos/gogate/apps/npis"
	"github.com/thinkgos/gogate/pkg/elink"
	"github.com/thinkgos/gogate/pkg/npi"
	"github.com/thinkgos/memlog"
)

// ZbNetworkController zigbee 网络控制器
type ZbNetworkController struct {
	ctrl.Controller
}

// Post 开启zigbee组网
func (this *ZbNetworkController) Post() {
	var err error
	var ok bool

	if npis.IsNetworkFormation() { // 开启组网
		ok, err = npis.ZbApps.Appcfg_BdbStartCommissioningReq(
			npi.Cms_mode_NetworkSteer)
	} else { // 建立网络并开启组网
		ok, err = npis.ZbApps.Appcfg_BdbStartCommissioningReq(
			npi.Cms_mode_NetworkFormation | npi.Cms_mode_NetworkSteer)
	}
	if err != nil || !ok {
		this.ErrorResponse(elink.CodeErrSysException)
		return
	}
	npis.SetNetworkSteering(true)
	memlog.Debug("elinkctls: zigbee network steering open")

	err = this.WriteResponsePyServerJSON(elink.CodeSuccess, nil)
	if err != nil {
		this.ErrorResponse(elink.CodeErrSysException)
		memlog.Error("response failed", err)
	}
}

// Delete 关闭zigbee组网
func (this *ZbNetworkController) Delete() {
	ok, err := npis.ZbApps.Zb_PermitJoingReq(0xfffc, 0)
	if err != nil || !ok {
		this.ErrorResponse(elink.CodeErrSysException)
		return
	}
	npis.SetNetworkSteering(false)
	memlog.Debug("elinkctls: zigbee network steering close")

	err = this.WriteResponsePyServerJSON(elink.CodeSuccess, nil)
	if err != nil {
		this.ErrorResponse(elink.CodeErrSysException)
		memlog.Error("response failed", err)
	}
}
