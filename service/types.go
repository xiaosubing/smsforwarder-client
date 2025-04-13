package service

// 基础信息
type PhoneBase struct {
	Code int           `json:"code"`
	Data PhoneBaseData `json:"data"`
}
type Num0 struct {
	Number string `json:"number"`
}
type Num1 struct {
	Number string `json:"number"`
}
type SimInfoList struct {
	Num0 Num0 `json:"0"`
	Num1 Num1 `json:"1"`
}
type PhoneBaseData struct {
	ExtraSim1   string      `json:"extra_sim1"`
	ExtraSim2   string      `json:"extra_sim2"`
	SimInfoList SimInfoList `json:"sim_info_list"`
}

type getMessageCode struct {
	Data      getMessageCodeData `json:"data"`
	Timestamp int64              `json:"timestamp"`
	Sign      string             `json:"sign"`
}
type getMessageCodeData struct {
	Type     int    `json:"type"`
	PageNum  int    `json:"page_num"`
	PageSize int    `json:"page_size"`
	Keyword  string `json:"keyword"`
}
