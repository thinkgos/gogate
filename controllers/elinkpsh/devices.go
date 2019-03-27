package elinkpsh

import (
	"github.com/slzm40/gogate/apps/mq"
	"github.com/slzm40/gogate/protocol/elinkres"
	"github.com/slzm40/gogate/protocol/elmodels"
	"github.com/slzm40/gomo/elink"

	"github.com/json-iterator/go"
)

// 设备加入或离开通知
func DeviceAnnce(pid int, sn string, isjoin bool) error {
	v, err := jsoniter.Marshal(elmodels.BaseSnPayload{pid, sn})
	if err != nil {
		return err
	}
	method := elink.MethodDelete
	if isjoin {
		method = elink.MethodPost
	}

	return mq.WriteCtrlData(elink.FormatResouce(elinkres.Devices, pid),
		method, elink.MessageTypeAnnce, v)
}
