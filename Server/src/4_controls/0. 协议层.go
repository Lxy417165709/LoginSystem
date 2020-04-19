package controls

//后端响应数据通信协议
type ResponseProto struct {
	Status int         `json:"status"` //状态 0正常，小于0出错，大于0可能有问题
	Msg    string      `json:"msg"`    //状态信息
	Data   interface{} `json:"data"`
}

// 前端请求数据通讯协议
type ReqProto struct {
	Data     interface{} `json:"data"`     //请求数据
	OrderBy  string      `json:"orderBy"`  //排序要求
	Filter   string      `json:"filter"`   //筛选条件
	Page     int         `json:"page"`     //分页
	PageSize int         `json:"pageSize"` //分页大小
}

// 修改用户个人信息数据结构
type UpiData struct {
	UserName         string `json:"userName"`
	UserSex          int    `json:"userSex"`
	UserBirthday     int    `json:"userBirthday"`
	UserContactPhone string `json:"userContactPhone"`
	UserContactEmail string `json:"userContactEmail"`
}

//// 对应数据库的 tb_userPersonalInformation 表
//type UserPersonalInformation struct {
//	UserId           int    `json:"userId" isPrimaryKey:"true"`
//	UserPhotoUrl     string `json:"userPhotoUrl"`
//	UserName         string `json:"userName"`
//	UserSex          int    `json:"userSex"`
//	UserContactPhone string `json:"userContactPhone"`
//	UserContactEmail string `json:"userContactEmail"`
//	UserBirthday     int    `json:"userBirthday"`
//	Reserved1        string `json:"reserved1"`
//	Reserved2        string `json:"reserved2"`
//}

// 注册需要三个字段 email,password,vrc
type RegisterData struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Vrc      string `json:"vrc"`
}

// 登录需要两个字段 email,password
type LoginData struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// 邮箱发送数据结构
type EmailData struct {
	Email string `json:"email"`
}

// 邮箱验证接口数据结构
type EmailVerificationData struct {
	Email string `json:"email"`
	Vrc   string `json:"vrc"`
}

// 修改用户头像数据结构
type UploadPhotoData struct {
	PhotoBase64 string `json:"photoBase64"`
}

// 获取图片接口数据结构
type GetPhotoData struct {
	PhotoName string `json:"photoName"`
}

// 图片响应数据结构
type GetPhotoRspData struct {
	PhotoBase64 string `json:"photoBase64"`
}

// 修改密码数据结构
type ChangePasswordData struct {
	NewPassword string `json:"newPassword"`
}
