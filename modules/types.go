package models

import "gorm.io/gorm"

type Message struct {
	gorm.Model
	ID      uint   `gorm:"primaryKey"`
	Phone   string `gorm:"size:20"`
	Number  string `gorm:"size:128"`
	Content string `gorm:"type:text"`
	Code    string `gorm:"size:10"`
}

// QueryParams 查询参数结构体
type QueryParams struct {
	Type     int    `json:"type"`      // 类型
	PageNum  int    `json:"page_num"`  // 页码
	PageSize int    `json:"page_size"` // 每页数量
	Keyword  string `json:"keyword"`   // 关键字
}
