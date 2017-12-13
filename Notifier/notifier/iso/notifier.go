package iso

import "fmt"
import (
	"github.com/NaySoftware/go-fcm"
	"github.com/Richardmbs12/Notifier/notifier"
)

type Notifier struct{
	ServerKey string
}


func (n *Notifier) Send(deviceId string, notification notifier.NotificationPayload, data map[string]string ){
	c := fcm.NewFcmClient(n.ServerKey)

	var notificationPayload = convert(notification)
	c.SetNotificationPayload(&notificationPayload)
	c.SetTimeToLive(4)
	c.SetPriority("High")
	c.NewFcmMsgTo(deviceId,data)

	status, err := c.Send()

	if err == nil {
		status.PrintResults()
	} else {
		fmt.Println(err)
	}
}

func convert(payload notifier.NotificationPayload) fcm.NotificationPayload {
	return fcm.NotificationPayload{
		Title: payload.Title,
		Body: payload.Body,
		Badge: payload.Badge,
		BodyLocArgs: payload.BodyLocArgs,
		AndroidChannelID: payload.AndroidChannelID,
		BodyLocKey: payload.BodyLocKey,
		ClickAction: payload.ClickAction,
		Color: payload.Color,
		Icon: payload.Icon,
		Sound: payload.Sound,
		Tag: payload.Tag,
		TitleLocArgs:payload.TitleLocArgs,
		TitleLocKey: payload.TitleLocKey,

	}
}
