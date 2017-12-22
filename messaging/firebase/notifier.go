package firebase

import (
	"github.com/NaySoftware/go-fcm"
	"github.com/fortresswallet/notifier/messaging"
)

type Notifier struct {
	ServerKey string
}

func (n *Notifier) SendBroadcast(request messaging.SendBroadcastRequest) error {
	c := fcm.NewFcmClient(n.ServerKey)

	var notificationPayload = convert(request.Payload)
	c.SetNotificationPayload(&notificationPayload)
	c.SetTimeToLive(4)
	c.SetPriority("High")
	c.NewFcmRegIdsMsg(request.List, request.Data)

	_, err := c.Send()

	return err
}

func (n *Notifier) Send(deviceId string, notification messaging.NotificationPayload, data map[string]string) error  {
	c := fcm.NewFcmClient(n.ServerKey)

	var notificationPayload = convert(notification)
	c.SetNotificationPayload(&notificationPayload)
	c.SetTimeToLive(4)
	c.SetPriority("High")
	c.NewFcmMsgTo(deviceId, data)

	_, err := c.Send()
	return err

}

func convert(payload messaging.NotificationPayload) fcm.NotificationPayload {
	return fcm.NotificationPayload{
		Title:            payload.Title,
		Body:             payload.Body,
		Badge:            payload.Badge,
		BodyLocArgs:      payload.BodyLocArgs,
		AndroidChannelID: payload.AndroidChannelID,
		BodyLocKey:       payload.BodyLocKey,
		ClickAction:      payload.ClickAction,
		Color:            payload.Color,
		Icon:             payload.Icon,
		Sound:            payload.Sound,
		Tag:              payload.Tag,
		TitleLocArgs:     payload.TitleLocArgs,
		TitleLocKey:      payload.TitleLocKey,
	}
}
