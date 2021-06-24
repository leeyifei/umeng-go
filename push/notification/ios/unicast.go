package ios

import (
	"encoding/json"
	"github.com/yzchan/umeng-go/push/notification"
	"time"
)

type Unicast struct {
	notification.Cast
	DeviceTokens string  `json:"device_tokens"`
	Payload      Payload `json:"payload"`
	Policy       Policy  `json:"policy,omitempty"`
	TemplateName string  `json:"template_name,omitempty"`
}

func NewUnicast() *Unicast {
	cast := &Unicast{}
	cast.Type = "unicast"
	cast.Payload = make(Payload)
	cast.Payload.Initial()
	cast.Timestamp = time.Now().Unix()
	return cast
}

func (cast *Unicast) SetDeviceToken(token string) *Unicast {
	if cast.TemplateName != "" { // 模板消息 禁止设置DeviceTokens
		return cast
	}
	cast.DeviceTokens = token
	return cast
}

func (cast *Unicast) Send() (string, error) {
	return cast.BaseSend(cast)
}

func (cast *Unicast) String() string {
	marshaled, _ := json.MarshalIndent(cast, "", "    ")
	return string(marshaled)
}

func (cast *Unicast) SetAsTemplate(name string) {
	cast.TemplateName = name
	cast.DeviceTokens = "${device_tokens}"
}
