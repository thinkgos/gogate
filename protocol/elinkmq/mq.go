package elinkmq

import (
	"github.com/thinkgos/gomo/elink"

	"github.com/astaxie/beego/logs"
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

var _ elink.Provider = (*Provider)(nil)

type Provider struct {
	C mqtt.Client
}

func NewProvider(c mqtt.Client) *Provider {
	return &Provider{c}
}

// 错误加在主题上的回复
func (this *Provider) ErrorDefaultResponse(topic string) error {
	return this.WriteResponse(topic, "{}")
}

// 应答信息
func (this *Provider) WriteResponse(topic string, data interface{}) error {
	return this.C.Publish(topic, 2, false, data).Error()
}

// 数据推送
func (this *Provider) Publish(tp string, data interface{}) error {
	return this.C.Publish(tp, 1, false, data).Error()
}

// 回调
func Handle(client mqtt.Client, message mqtt.Message) {
	logs.Debug("Topic: %s", message.Topic())
	logs.Warn("MessageID: %d,Qos - %d,Retained - %t,Duplicate - %t",
		message.MessageID(), message.Qos(), message.Retained(), message.Duplicate())
	//logs.Debug("receive:\n%s\n", message.Payload())

	// 抛弃retain 和重复的消息 必须使用Qos = 2的消息
	if message.Retained() || message.Duplicate() || message.Qos() != 2 {
		logs.Warn("Handle: Invalid message discard")
		return
	}
	elink.Server(NewProvider(client), message.Topic(), message.Payload())
}