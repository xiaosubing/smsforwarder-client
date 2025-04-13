package conf

import (
	"bufio"
	"fmt"
	"github.com/spf13/viper"
	"os"
)

var Smsforwarder *smsforwarderSetting
var Message = make(chan string, 5)

type smsforwarderSetting struct {
	BaseInfo         *baseinfo
	MessageTemplate  string
	GetMessageVerify string
	GetEncryptPhone  string
	Db               *saveMessage
	Notify           *notify
}
type baseinfo struct {
	PhoneNumber string
	PhoneAlias  string
}

type saveMessage struct {
	SaveType    string
	DbType      string
	DbName      string
	Encrypt     bool
	EncryptSalt string
	DbUser      string
	DbPassword  string
	DbHost      string
}
type notify struct {
	NotifyType           []string
	NotifyWebHookUrl     string
	NotifyWebHookType    string
	NotifyWebHookPayload string

	NotifyMailAccount  string
	NotifyMailPassword string
	NotifyMailSendTo   string
	NotifyMailSmtpHost string
	NotifyMailSmtpPort int
	NotifyMailSubject  string

	CodeSecON   bool
	CodeSecType string
	// more notify ...
}

func init() {
	viper.SetConfigName("conf")
	viper.SetConfigType("yml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("./")

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("未找到配置，已生成默认配置文件， 请编辑url后重新运行！！！")
		createConf()
		os.Exit(0)
	}

	Smsforwarder = NewSmsforwarder()

	if len(Smsforwarder.BaseInfo.PhoneNumber) == 0 {
		fmt.Println("请配置好基础信息,然后重新启动")
		os.Exit(1)
	}
	if len(Smsforwarder.BaseInfo.PhoneAlias) == 0 {
		Smsforwarder.BaseInfo.PhoneAlias = Smsforwarder.BaseInfo.PhoneNumber
	}

}

func NewSmsforwarder() *smsforwarderSetting {

	return &smsforwarderSetting{
		BaseInfo: &baseinfo{
			PhoneNumber: viper.GetString("baseInfo.phoneNumber"),
			PhoneAlias:  viper.GetString("baseInfo.phoneAlias"),
		},
		MessageTemplate:  viper.GetString("template"),
		GetMessageVerify: viper.GetString("getMessage.verify"),

		Db: &saveMessage{
			SaveType:    viper.GetString("db.savetype"),
			DbType:      viper.GetString("db.dbtype"),
			DbName:      viper.GetString("db.name"),
			Encrypt:     viper.GetBool("db.encrypt"),
			EncryptSalt: viper.GetString("db.encryptSalt"),
			DbUser:      viper.GetString("db.user"),
			DbPassword:  viper.GetString("db.password"),
			DbHost:      viper.GetString("db.host"),
		},

		Notify: &notify{
			NotifyType:           viper.GetStringSlice("notify.type"),
			NotifyWebHookUrl:     viper.GetString("notify.webhook.url"),
			NotifyWebHookType:    viper.GetString("notify.webhook.type"),
			NotifyWebHookPayload: viper.GetString("notify.webhook.payload"),
			NotifyMailAccount:    viper.GetString("notify.mail.account"),
			NotifyMailPassword:   viper.GetString("notify.mail.password"),
			NotifyMailSendTo:     viper.GetString("notify.mail.sendTo"),
			NotifyMailSmtpHost:   viper.GetString("notify.mail.smtpHost"),
			NotifyMailSmtpPort:   viper.GetInt("notify.mail.smtpPort"),
			NotifyMailSubject:    viper.GetString("notify.mail.subject"),
			CodeSecON:            viper.GetBool("notify.codeSec.on"),
			CodeSecType:          viper.GetString("notify.codeSec.type"),
		},
	}

}
func createConf() {
	file, err := os.OpenFile("./conf.yml", os.O_CREATE|os.O_WRONLY, 0766)
	if err != nil {
		fmt.Println("创建配置文件失败，请手动创建")
		os.Exit(0)
	}
	defer file.Close()

	s := `
# 消息模板

# 基础信息
baseInfo:
  phoneNumber: 手机号，必须填写！
  phoneAlias: 手机号别名，请填写！

# 1 需要验证码  顺序不能乱！
template: "验证码: [验证码]\n收信人: [收信人]\n发信人: [发信人]\n短信原文:\n[短信原文]"

# getMessage 验证
getMessage:
  verify: ""
  encrypted: ""


# 消息保存地址
db:
  # 保存到本地还是远程
  savetype: local
  dbType: sqlite
  name: "/opt/smsforwarder/smsforwarder.db"
  # 是否加密
  encrypt: false
  # salt
  encryptSalt: "梅干菜小酥饼"
  # 远程地址
  host: ""
  # 用户名
  user: ""
  # 密码
  password: ""




# 配置通知渠道
notify:
  # 通知渠道，必填！！！！！！！！！！！！！！！！！！！！！！！！！
  # 可以配置的值为： qq 、webhook 和 mail
  # 填写完成后请完善对应渠道的详细信息！！！
  # 支持多渠道消息通知
  type:
    - webhook
    #- qq

  # 如果识别出了验证码是否再发送到其他渠道（单个)
  # 默认关闭
  codeSec:
    on: false
    type: mail

  webhook:
    url: "https://send-notifyme.521933.xyz"
    type: "post"
    payload: '{
      "data": {
        "to": "请输入你的token",
        "ttl": 86400,
        "priority": "normal",
        "pushType": "FCM_Push",
        "data": {
          "title": "[验证码]",
          "body": "[短信原文]",
          "group": "Messages",
          "bigText": false,
          "iconType": 0,
          "smallIcon":"smallIcon_0",
          "largeIcon":"largeIcon_0",
          "id":"0",
          "encryption":false,
          "iv":"UkAjUPykxX1Eu4h7",
          "invisible":false,
          "actionType":"0",
          "action":"",
          "channel":"quicker_channel",
          "record":1
        }
      }
    }'

    # 钉钉的payload, "[短信原文]" 表示短信内容，"[验证码]" 主动识别到的验证码，有可能有误识别，请注意！
    #payload: '{"msgtype": "text","text": {"content": "[短信原文]"}}'
    # 其他post请求的payload， 根据自己的请求配置
    #payload: '{"message": "[短信原文]", "to": "梅干菜小酥饼"}'


  # 邮箱通知
  mail:
    account: "3286276407@qq.com"
    password: "非邮箱登陆密码，请自行获取邮箱的凭证"
    sendTo: "建议使用运营商邮箱为收件人，打开短信提醒，能够自动识别验证码"

    subject: "[验证码]"
    # 其他邮箱可以把程序内自动识别的验证码加上，但是识别有可能不准哦
    # subject: "验证码：1"

    # qq邮箱的服务器信息
    # 默认使用qq邮箱发送, 可自行替换其他邮箱
    smtpHost: "smtp.qq.com"
    smtpPort: 587



`
	write := bufio.NewWriter(file)
	_, err = write.WriteString(s)
	if err != nil {
		fmt.Println("写入配置失败，请手动创建并配置配置文件")
		os.Exit(0)
	}

	write.Flush()
}
