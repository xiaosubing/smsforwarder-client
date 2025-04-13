package message

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/godbus/dbus/v5"
	"smsforwarder/conf"
	models "smsforwarder/modules"
	"smsforwarder/service"
)

type sendMessageReq struct {
	Data      Data   `json:"data"`
	Timestamp int64  `json:"timestamp"`
	Sign      string `json:"sign"`
}
type Data struct {
	SimSlot      int    `json:"sim_slot"`
	PhoneNumbers string `json:"phone_numbers"`
	MsgContent   string `json:"msg_content"`
}

func processMessage(text, sender string) {
	code := service.MessageCodeProcess(text)
	conf.Message <- fmt.Sprintf("%s---%s---%s", text, sender, code)
	fmt.Println("获取到的消息：", text)

	models.SaveMessage(text, sender, code)
	fmt.Println("消息发送完成")
}

func SendMessage(c *gin.Context) {
	var request = new(sendMessageReq)
	err := c.BindJSON(request)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Invalid request body: " + err.Error(),
		})
		return
	}

	//fmt.Println("===============执行短信发送任务==============")
	//fmt.Printf("发给： %s\n", request.Data.PhoneNumbers)
	//fmt.Printf("发送内容： %s\n", request.Data.MsgContent)
	//fmt.Println("==========================================")

	// Create SMS message
	messagingObj := conn.Object("org.freedesktop.ModemManager1", dbus.ObjectPath("/org/freedesktop/ModemManager1/Modem/0"))
	smsProps := map[string]dbus.Variant{
		"number": dbus.MakeVariant(request.Data.PhoneNumbers),
		"text":   dbus.MakeVariant(request.Data.MsgContent),
	}

	var smsPath dbus.ObjectPath
	err = messagingObj.Call("org.freedesktop.ModemManager1.Modem.Messaging.Create", 0, smsProps).Store(&smsPath)
	if err != nil {
		c.JSON(500, gin.H{
			"error": fmt.Sprintf("Failed to create SMS message: %v", err),
		})
		return
	}

	// Send SMS
	smsObj := conn.Object("org.freedesktop.ModemManager1", smsPath)
	err = smsObj.Call("org.freedesktop.ModemManager1.Sms.Send", 0).Err
	if err != nil {
		c.JSON(500, gin.H{
			"error": fmt.Sprintf("Failed to send SMS: %v", err),
		})
		return
	}

	// Return success response
	c.JSON(200, gin.H{
		"code":    200,
		"status":  "success",
		"message": "短信发送成功！",
	})
}
