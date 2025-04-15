package service

import (
	"encoding/json"
	"smsforwarder/conf"
	"strings"
)

func Forwarder() {
	for v := range conf.ForwarderMessage {

		if conf.Smsforwarder.Forwarder.ForwarderOn == true {
			text := strings.Split(v, "---")

			type RequestBody struct {
				Phone   string `json:"phone"`
				Number  string `json:"number"`
				Message string `json:"message"`
				Sign    string `json:"sign"`
			}

			payload, _ := json.Marshal(RequestBody{
				Phone:   conf.Smsforwarder.BaseInfo.PhoneNumber,
				Number:  text[1],
				Message: text[0],
				Sign:    "梅干菜小酥饼",
			})
			// forwarder
			HttpPost(conf.Smsforwarder.Forwarder.ForwarderUrl, string(payload))

		}

	}
}
