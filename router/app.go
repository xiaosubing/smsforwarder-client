package router

import (
	"github.com/gin-gonic/gin"
	"smsforwarder/message"
	"smsforwarder/service"
)

func App() *gin.Engine {
	r := gin.Default()

	r.POST("/sms/send", message.SendMessage)
	r.POST("/sms/query", service.GetMessageCode)
	r.POST("/config/query", service.GetBaseInfo)
	r.POST("/api/cmd", service.TodoCMD)

	return r
}
