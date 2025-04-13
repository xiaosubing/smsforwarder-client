package message

import (
	"fmt"
	"github.com/godbus/dbus/v5"
	"os"
	"time"
)

// vars
var (
	c       = make(chan *dbus.Signal, 2)
	conn, _ = dbus.SystemBus()
	rule    = "type='signal',interface='org.freedesktop.ModemManager1.Modem.Messaging',member='Added'"
)

type Response struct {
	Message string `json:"phone"`
}

func ListenMessage() {
	call := conn.BusObject().Call("org.freedesktop.DBus.AddMatch", 0, rule)
	if call.Err != nil {
		fmt.Println("Failed to add match: %v", call.Err)
		os.Exit(1)
	}
	conn.Signal(c)
	defer conn.Close()

	var tmp string
	for v := range c {
		if v.Body[1] == true {
			var text string
			var sender string
			service := conn.Object("org.freedesktop.ModemManager1", dbus.ObjectPath(fmt.Sprintf("%s", v.Body[0])))
			service.Call("org.freedesktop.DBus.Properties.Get", 0, "org.freedesktop.ModemManager1.Sms", "Text").Store(&text)
			service.Call("org.freedesktop.DBus.Properties.Get", 0, "org.freedesktop.ModemManager1.Sms", "Number").Store(&sender)
			if len(text) == 0 {
				for i := 0; i <= 20; i++ {
					time.Sleep(1 * time.Second)
					service.Call("org.freedesktop.DBus.Properties.Get", 0, "org.freedesktop.ModemManager1.Sms", "Text").Store(&text)
					service.Call("org.freedesktop.DBus.Properties.Get", 0, "org.freedesktop.ModemManager1.Sms", "Number").Store(&sender)
					if len(text) != 0 {
						break
					}
				}
			}

			if text != tmp {
				processMessage(text, sender)
				tmp = text
			}
		}

	}

}
