package commonConst

// 全局常量
const (
	ErrorFlag              = 0
	LoginSuccessFlag       = 1
	RegisterSuccessFlag    = 1
	UpdateSuccessFlag      = 1
	QuerySuccessFlag       = 1
	EmailSendSuccessFlag   = 1
	EmailVerifySuccessFlag = 1
	UpiUpdateSuccessFlag   = 1
	PhotoUploadSuccessFlag = 1
	GetPhotoSuccessFlag    = 1
	LinkValidFlag          = 1
	PasswordChangeFlag     = 1
	UaiGetSuccessFlag      = 1
	UpiGetSuccessFlag      = 1

	AESKey   = "ajwiskdlg129/452"
	TokenKey = "woshilixueyue"
	ConfPath = "src/conf.ini"
)

// 发送验证码相关
const (
	ChangePasswordExpiredTime = 300
	RegisterExpiredTime       = 60
)

// 用户相关常量
const (
	ErrorUserId      = 0
	DefaultPhotoUrl  = "test.gif"
	DefaultUserName  = "无名氏"
	DefaultUserSex   = Man
	SmallUser        = 1
	DefaultUserPhone = "18946910438"
	BirthDayRato     = 1000
)

// 盐值相关
const (
	SaltLength = 5
	SaltPool   = "abcdefghijklmnopqrstuvwsyz0123456789"
)

// 验证码相关
const (
	VrcLength = 6
	VrcPool   = "0123456789"
)

// 正则相关
// 正则可能有些错误
const (
	EmailRegexp    = `^[a-zA-Z0-9_.-]+@[a-zA-Z0-9-]+(\.[a-zA-Z0-9-]+)*\.[a-zA-Z0-9]{2,6}$`
	PasswordRegexp = `^[_.a-zA-Z0-9]{8,15}$`
	Base64Regexp   = ``
	PhoneRegexp    = `^((13[0-9])|(14[5,7])|(15[0-3,5-9])|(17[0,3,5-8])|(18[0-9])|166|198|199|(147))\d{8}$`
	PhotoRegexp    = `^image/.+`
)

// 单位相关
const (
	_  = iota
	KB = 1 << (iota * 10)
	MB
)

//type sex = int
const (
	_ int = iota
	Man
	Woman
)

// 校验器相关
const(
	RegisterVrcKeyPrefix = "registerVrc"
	ChangePasswordVrcKeyPrefix  = "changePasswordVrc"
)

// 表名
const (
	NameOfTableUai = "tb_userAccountInformation"
	NameOfTableUpi = "tb_userPersonalInformation"
)
