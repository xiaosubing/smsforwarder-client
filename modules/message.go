package models

import (
	"fmt"
	"smsforwarder/conf"
)

func SaveMessage(text, sender, code string) {
	var m Message
	m.Phone = conf.Smsforwarder.BaseInfo.PhoneNumber
	m.Number = sender
	m.Code = code
	m.Content = text

	err := InsertData(&m)
	if err != nil {
		return
	}
}

func GetMessages(param QueryParams) []Message {
	var m []Message

	err := QueryData(&m, param)
	if err != nil {
		fmt.Println("执行出错拉！ ")
		fmt.Println(err)
		return nil
	}

	return m
}
