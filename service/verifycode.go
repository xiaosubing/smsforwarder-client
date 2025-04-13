package service

import (
	"regexp"
	"strings"
)

// MessageCodeProcess 正则获取验证码
// 【优酷土豆】您的短信验证码是811245。您的手机号正在使用随机密码登录服务，如非本人操作，请尽快修改密码。
// 以下类型验证码获取存在问题， 待修复
// 991378(登录随机码) ，感谢您使用中国联通APP【中国联通】               来源： wap.10010.com
// 【芒果tv】338673（随机验证码），有效期10分钟。如非本人使用，敬请忽略本信息。
// 【沃畅玩】请勿将验证码告知任何人，您的验证码为：316235，有效时间5分钟。
func MessageCodeProcess(content string) string {
	re := regexp.MustCompile(`.*?(.{0,15})[随机|验证|登录|授权|动态|校验]码(.{0,10})`)
	match := re.FindAllString(content, -1)
	if len(match) == 0 {
		return "None"
	}
	re = regexp.MustCompile(`\d{4,6}\b`)
	code := re.FindAllString(match[0], -1)

	if len(code) == 0 {
		content = "验证" + strings.Replace(content, match[0], "", -1)
		//re = regexp.MustCompile(`.*?(.{0,15})[随机|验证|登录|授权|动态|校验]码(.{0,10})`)
		match = re.FindAllString(content, -1)
		if len(match) == 0 {
			return "None"
		}
		re = regexp.MustCompile(`\d{4,6}\b`)
		code = re.FindAllString(match[0], -1)
	}

	if len(code) != 0 {
		return code[0]
	}
	return "None"
}
