package internal

import (
	"fmt"

	"github.com/esiqveland/notify"
	"github.com/godbus/dbus"
	"github.com/skratchdot/open-golang/open"
)

func Notify(level Level, timeout int, title, message, detailFilename string) error {
	conn, err := dbus.SessionBus()
	if err != nil {
		return fmt.Errorf("failed to connect session BUS %w", err)
	}

	buttons := []string{"ok", "OK"} // tuples of (action_key, label)
	if detailFilename != "" {
		buttons = []string{"details", "Details", "ok", "OK"}
	}

	// Create a Notification to send
	n := notify.Notification{
		ReplacesID:    uint32(0),
		AppIcon:       level.String(),
		Summary:       title,
		Body:          message,
		Actions:       buttons,
		ExpireTimeout: int32(timeout * 1000),
	}

	// Notifyer interface with event delivery
	notifier, err := notify.New(conn)
	if err != nil {
		return fmt.Errorf("failed to build notifier %w", err)
	}
	defer notifier.Close()

	if _, err := notifier.SendNotification(n); err != nil {
		return fmt.Errorf("error sending notification: %w", err)
	}

	// Listen for actions invoked!
	actions := notifier.ActionInvoked()
	go func() {
		action := <-actions
		switch action.ActionKey {
		case "details":
			if err := open.Start(detailFilename); err != nil {
				panic(err)
			}
		}
	}()

	<-notifier.NotificationClosed()

	return nil
}
