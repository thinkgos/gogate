package mq

import (
	"fmt"
	"sync"
	"time"

	"github.com/slzm40/gogate/protocol/elmodels"
	"github.com/slzm40/gomo/elink"
	"github.com/slzm40/gomo/misc"
	"github.com/slzm40/gomo/protocol/elinkch/ctrl"

	"github.com/astaxie/beego/logs"
	"github.com/eclipse/paho.mqtt.golang"
	"github.com/json-iterator/go"
)

const (
	mqtt_broker_address  = "mqtt.lchtime.cn:1883"
	mqtt_broker_password = "52399399"
)

const (
	gatewayProductKey = "lc_gzs100"
)

var gatewayHeartBeatTopic = fmt.Sprintf("data/0/%s/gateway.heartbeat/patch/time", misc.Mac())
var Client mqtt.Client
var heartOnce sync.Once

func init() {
	elink.RegisterTopicInfo(misc.Mac(), gatewayProductKey) // 注册网关产品Key

	opts := mqtt.NewClientOptions()
	opts.AddBroker(mqtt_broker_address).SetClientID(misc.Mac()) // broker and clientID
	opts.SetUsername("1").SetPassword(mqtt_broker_password)     // user name and password
	opts.SetCleanSession(false).SetAutoReconnect(true)

	opts.SetOnConnectHandler(func(cli mqtt.Client) {
		logs.Info("mqtt client connect success")
		chList := elink.ChannelSelectorList()
		for _, ch := range chList {
			s := fmt.Sprintf("%s/%s/%s/+/+/+/#", ch, elink.TpInfos.ProductKey, misc.Mac())
			cli.Subscribe(s, 2, elink.Server)
		}
		heartOnce.Do(func() {
			time.AfterFunc(time.Second, HeartBeatStatus)
			fmt.Println("once???????????????????????")
		})
	})

	opts.SetConnectionLostHandler(func(cli mqtt.Client, err error) {
		logs.Warn("mqtt client connection lost, ", err)
	})

	if out, err := jsoniter.Marshal(elmodels.GatewayHeatbeats(false)); err != nil {
		logs.Error("mqtt %s", err.Error())
	} else {
		opts.SetBinaryWill(gatewayHeartBeatTopic, out, 2, false)
	}
	Client = mqtt.NewClient(opts)
	started()
}

// 启动连接mqtt
func started() {
	logs.Info("mqtt client connecting...")
	if token := Client.Connect(); token.Wait() && token.Error() != nil {
		logs.Error("mqtt client connect failed, ", token.Error())
		time.AfterFunc(time.Second*30, started)
	}
}

// 网关心跳包
func HeartBeatStatus() {
	defer time.AfterFunc(time.Second*30, HeartBeatStatus)
	if !Client.IsConnected() {
		return
	}

	out, err := jsoniter.Marshal(elmodels.GatewayHeatbeats(true))
	if err != nil {
		logs.Error("HeartBeatStatus:", err)
		return
	}
	Client.Publish(gatewayHeartBeatTopic, 0, false, out)

}

// ctrl data通道推送数据
func WriteCtrlData(resourse, method, messageType string, payload []byte) error {
	return ctrl.WriteData(Client, resourse, method, messageType, payload)
}
