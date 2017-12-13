package main

import (
	"github.com/Richardmbs12/notifier/notifier/iso"
	"github.com/Richardmbs12/notifier/notifier"
)

func main() {
	//Instantiate a notifier
	isoNotifier := iso.Notifier{
		ServerKey:"<ServerKey>",
	}

	n := NotifierContract{
		Notifier: &isoNotifier,
	}

	var notification notifier.NotificationPayload
	notification.Body = "Hello World, this is content body"
	notification.Title = "This is a title!"

	n.Notifier.Send("<DeviceID>",notification,map[string]string{
		"msg": "Hello World1",
		"sum": "Happy Day",
	} )

}

type NotifierContract struct {
	Notifier notifier.Notifier
}
