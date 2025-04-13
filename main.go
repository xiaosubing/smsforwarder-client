package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"smsforwarder/message"
	models "smsforwarder/modules"
	"smsforwarder/notify"
	"smsforwarder/router"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	gin.DefaultWriter = ioutil.Discard

	go models.NewDB()
	go message.ListenMessage()
	go notify.Notify()

	r := router.App()

	//for i := 0; i < 10; i++ {
	//	time.Sleep(time.Second * 1)
	//	models.SaveMessage("测试", "10010", "123456")
	//}
	fmt.Println("开始监听短信......")

	r.Run(":5000")
	select {}
}
