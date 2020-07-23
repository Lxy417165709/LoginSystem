package middleware

import (
	"fmt"
	"github.com/astaxie/beego/logs"
	"4_controls/interaction"
	"net/http"
	"time"
)
type functionType = func (w http.ResponseWriter,r *http.Request)

var counterMap = make(map[string]int)	// 并发不安全
var interactionManager = interaction.GetInteractionManger()
const (
	forbitUserAgent = "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/71.0.3578.98 Safari/537.36"
)
// 计算操作耗时的中间件
func TimeCounter(function functionType) functionType{
	return func (w http.ResponseWriter,r *http.Request) {
		begin:= time.Now()
		function(w,r)
		end :=time.Since(begin)
		logs.Info("%s 接口: 访问时间为 %d 微秒",r.URL.String(),end.Microseconds())
	}
}

// 计算访问次数的中间件
func TimesCounter(function functionType) functionType{
	return func (w http.ResponseWriter,r *http.Request) {
		function(w,r)
		counterMap[r.URL.String()]++
		logs.Info("%s 接口: 访问次数为 %d",r.URL.String(),counterMap[r.URL.String()])
	}
}

// 限制浏览器访问的中间件
func Limit(function functionType) functionType{
	return func (w http.ResponseWriter,r *http.Request) {
		if r.UserAgent()==forbitUserAgent{
			interactionManager.ResponseError(w,fmt.Errorf("请换种浏览方式~"))
			return
		}
		//fmt.Println(r.UserAgent())
		function(w,r)
	}
}


// 防止用户重复登录的中间件(这个成功了) 不过有点乱，需要借助tokenMap结构
//func ForbidRepeatLogin(function functionType) functionType{
//	return func (w http.ResponseWriter,r *http.Request) {
//		var uid int
//		var Err *commonStruct.Error
//		if uid,Err = controls.GetUidFromRequest(r);Err!=nil{
//			controls.HandleErr(w,Err)
//			return
//		}
//		tk,err := controls.GetTokenFromRequest(r)
//		if err!=nil{
//			logs.Error(err)
//			return
//		}
//		if controls.TokenMap[uid] != tk{
//			controls.ResponseError(w,fmt.Errorf("登录信息失效"))
//		}else{
//			function(w,r)
//		}
//
//	}
//}
