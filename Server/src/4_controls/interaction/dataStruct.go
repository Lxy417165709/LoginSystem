package interaction

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

// 登录需要两个字段 email,password
type LoginDTO struct {
	Email    string `json:"email" checkType:"email"`
	Password string `json:"password" checkType:"password"`
}


// 注册需要三个字段 email,password,vrc
type RegisterDTO struct {
	Email    string `json:"email" checkType:"email"`
	Password string `json:"password" checkType:"password"`
	Vrc      string `json:"vrc" checkType:"vrc"`
}

// 更新用户信息的数据结构
type UpiDTO struct {
	UserName         string `json:"userName" checkType:"userName"`
	UserSex          int    `json:"userSex" checkType:"sex"`
	UserContactPhone string `json:"userContactPhone" checkType:"phone"`
	UserContactEmail string `json:"userContactEmail" checkType:"email"`
	UserBirthday     int    `json:"userBirthday" checkType:"birthday"`
}

// 修改用户头像数据结构
type UpdatePhotoDTO struct {
	PhotoBase64 string `json:"photoBase64" checkType:"photo"`
}

// 邮箱发送验证码数据结构
type EmailDTO struct {
	Email string `json:"email" checkType:"email"`
}

// 邮箱验证接口数据结构
type EmailVerificationDTO struct {
	Email string `json:"email"`
	Vrc   string `json:"vrc"`
}



// 获取图片接口数据结构
type GetPhotoDTO struct {
	PhotoName string `json:"photoName"`
}

// 图片响应数据结构
type GetPhotoRspDTO struct {
	PhotoBase64 string `json:"photoBase64"`
}

// 修改密码数据结构
type ChangePasswordDTO struct {
	NewPassword string `json:"newPassword"`
}


