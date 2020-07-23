package checker

import "3_transition"

type checkResult struct {
	ForDeveloper error
	ForUser      error
}

var checkerInstance = &Checker{make(map[int]string)}
var fpi = transition.GetFPI()

type Checker struct {
	msgMap map[int]string
}

func (c *Checker) init() {
	c.msgMap[ErrorFlag] = "服务器在校验表单时发生错误"
	c.msgMap[EmailNotExistFlag] = "邮箱不存在"
	c.msgMap[PasswordNotRightFlag] = "密码错误"
	c.msgMap[VrcEmptyFlag] = "验证码未发送或已过期"
	c.msgMap[VrcErrorFlag] = "验证码错误"
	c.msgMap[EmailFormatErrorFlag] = "邮箱格式错误"
	c.msgMap[PasswordFormatErrorFlag] = "密码格式错误"
	c.msgMap[VrcFormatErrorFlag] = "验证码格式错误"
	c.msgMap[PhoneFormatErrorFlag] = "手机格式错误"
	c.msgMap[SexSelectErrorFlag] = "性别选择错误"
	c.msgMap[BirthdaySelectErrorFlag] = "生日选择错误"
	c.msgMap[PhotoIsNotPhotoFileFlag] = "该文件不是图片文件"
	c.msgMap[PhotoTooLargeFlag] = "文件太大了"
	c.msgMap[UsernameFormatErrorFlag] = "用户名格式错误"
	c.msgMap[EmailExistErrorFlag] = "该邮箱已存在"
}

func GetChecker() *Checker {
	checkerInstance.init()
	return checkerInstance
}
