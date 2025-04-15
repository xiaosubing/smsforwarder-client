package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"smsforwarder/message"
	models "smsforwarder/modules"
	"smsforwarder/notify"
	"smsforwarder/router"
	"smsforwarder/service"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	gin.DefaultWriter = ioutil.Discard

	go models.NewDB()
	go message.ListenMessage()
	go notify.Notify()
	go service.Forwarder()

	r := router.App()

	fmt.Println("开始监听短信......")

	r.Run(":5000")
	select {}
}
