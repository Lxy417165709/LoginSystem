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

