package main

import (
	"0_common/commonConst"
	"1_env"
	"3_transition"
	"5_middleware"
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/gorilla/websocket"
	"net/http"
	"time"
)

func AllInit() {
	// 初始化日志器
	logs.SetLogFuncCallDepth(3)
	logs.EnableFuncCallDepth(true)

	// 初始化配置文件
	env.LoadConf(commonConst.ConfPath)

	// 初始化数据库、校验器
	if err := transition.Init(); err != nil {
		logs.Info("初始化失败: " + err.Error())
	}
}

func AllClose() {
	transition.Close()
}

func main() {
	AllInit()
	defer func() {
		AllClose()
	}()

	logs.Info("init successfully!!!")

	http.HandleFunc("/server/test", middleware.Test)

	// 不用表单
	http.HandleFunc("/server/getPhoto", middleware.GetPhoto)
	http.HandleFunc("/server/getUai", middleware.GetUai)
	http.HandleFunc("/server/getUpi", middleware.GetUpi)

	// 要表单
	http.HandleFunc("/server/login", middleware.Login)
	http.HandleFunc("/server/register", middleware.Register)
	http.HandleFunc("/server/updateUserPersonalInformation", middleware.UpdateUpi)
	http.HandleFunc("/server/updatePhoto", middleware.UpdatePhoto)
	http.HandleFunc("/server/registerVrc/send", middleware.SendRegisterVrc)

	// 测试websocket
	http.HandleFunc("/server/wbsk", wbsk)

	if err := http.ListenAndServe(fmt.Sprintf(":%d", env.Conf.Server.Port), nil); err != nil {
		logs.Error(err)
		return
	}
}

func wbsk(w http.ResponseWriter, r *http.Request) {
	conn, err := (&websocket.Upgrader{CheckOrigin: func(r *http.Request) bool { return true }}).Upgrade(w, r, nil)
	if err != nil {
		panic(err)
	}
	for {
		if err := conn.WriteMessage(websocket.TextMessage, []byte("hello")); err != nil {
			panic(err)
		}
		time.Sleep(5 * time.Second)
	}
}
