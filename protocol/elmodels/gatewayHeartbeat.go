package elmodels

import (
	"time"

	"github.com/thinkgos/common"
	"github.com/thinkgos/gomo/misc"
)

type DeviceInfo struct {
	Sn string `json:"sn"`
}
type DeviceStatus struct {
	CurrentTime   string `json:"currentTime"`
	StartDateTime string `json:"startDateTime"`
	RunningTime   string `json:"runningTime"`
	Status        string `json:"status"`
}
type NetInfo struct {
	MAC string `json:"MAC"`
	Mac string `json:"mac"`
}

type Info struct {
	Uid          []int        `json:"uid"`
	DeviceInfo   DeviceInfo   `json:"device_info"`
	DeviceStatus DeviceStatus `json:"device_status"`
	NetInfo      NetInfo      `json:"net_info"`
}

type GatewayHeatbeat struct {
	Info Info `json:"info"`
}

func GatewayHeatbeats(isonline bool) *GatewayHeatbeat {
	status := "online"
	if !isonline {
		status = "offline"
	}
	mac := misc.Mac()
	return &GatewayHeatbeat{
		Info: Info{
			DeviceInfo: DeviceInfo{Sn: mac},
			DeviceStatus: DeviceStatus{
				CurrentTime:   time.Now().Local().Format("2006-01-02 15:04:05"),
				StartDateTime: common.SetupTime(),
				RunningTime:   common.RunningTime(),
				Status:        status,
			},
			NetInfo: NetInfo{
				MAC: misc.MAC(),
				Mac: mac,
			},
		},
	}
}
