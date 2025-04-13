package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"smsforwarder/conf"
	models "smsforwarder/modules"
)

func GetBaseInfo(c *gin.Context) {

	var devices PhoneBase

	devices.Code = 200
	devices.Data.ExtraSim1 = conf.Smsforwarder.BaseInfo.PhoneAlias
	devices.Data.SimInfoList.Num0.Number = conf.Smsforwarder.BaseInfo.PhoneNumber

	c.JSON(200, devices)

}

func GetMessageCode(c *gin.Context) {
	recInfo := new(getMessageCode)
	err := c.ShouldBind(recInfo)
	if err != nil {
		c.JSON(200, gin.H{
			"code": -1,
			"msg":  404,
		})
		return
	}
	var param models.QueryParams
	param.PageSize = recInfo.Data.PageSize
	if recInfo.Data.Keyword != "" {
		param.Keyword = fmt.Sprintf("content LIKE \"%s%s%s\"", "%", recInfo.Data.Keyword, "%")
	}

	msg := models.GetMessages(param)
	//fmt.Println(len(msg))
	for _, msg := range msg {
		fmt.Println(msg)
	}
	c.JSON(200, msg)

}
