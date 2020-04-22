package middleware

import (
	"github.com/astaxie/beego/logs"
	"net/http"
	"time"
)
type functionType = func (w http.ResponseWriter,r *http.Request)

var counterMap = make(map[string]int)	// 并发不安全

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
