package main

import (
	"github.com/thinkgos/gogate/apps/mq"
	"github.com/thinkgos/gogate/apps/npis"
	"github.com/thinkgos/gogate/misc"
	"github.com/thinkgos/gogate/misc/discover"

	"github.com/astaxie/beego"

	_ "github.com/thinkgos/gogate/models"
	_ "github.com/thinkgos/gogate/routers"
)

func main() {
	misc.LogsInit()
	mq.MqttInit()
	if npis.ZbAppInit() != nil {
		panic("main: npi app init failed")
	}
	//logs.Debug("here")

	go discover.Run("192.168.199.240")

	beego.Run()
}
