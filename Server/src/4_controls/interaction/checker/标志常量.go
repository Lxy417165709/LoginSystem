package checker

// 错误的标识符 (偶数)
const (
	ErrorFlag         = 0
	EmailNotExistFlag = 1 << iota
	PasswordNotRightFlag
	VrcErrorFlag
	VrcEmptyFlag
	EmailFormatErrorFlag
	PasswordFormatErrorFlag
	VrcFormatErrorFlag
	PhoneFormatErrorFlag
	SexSelectErrorFlag
	BirthdaySelectErrorFlag
	PhotoIsNotPhotoFileFlag
	PhotoTooLargeFlag
	UsernameFormatErrorFlag
	EmailExistErrorFlag
)

// 正确的标识符 (奇数)
const (
	_              = iota
	EmailExistFlag = (1 << iota) - 1
	EmailExistRightFlag
	PasswordRightFlag
	VrcRightFlag
	EmailFormatRightFlag
	PasswordFormatRightFlag
	VrcFormatRightFlag
	PhoneFormatRightFlag
	SexSelectRightFlag
	BirthdaySelectRightFlag
	PhotoValidFlag
	UsernameFormatRightFlag
	CheckPassFlag
)
