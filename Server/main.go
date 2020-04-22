package main

import (
	"0_common/commonConst"
	"1_env"
	"3_transition"
	"5_middleware"
	"fmt"
	"github.com/astaxie/beego/logs"
	"net/http"
)

func AllInit() error {
	// 初始化日志器
	logs.SetLogFuncCallDepth(3)
	logs.EnableFuncCallDepth(true)

	// 初始化配置文件
	if err := env.LoadConf(commonConst.ConfPath); err != nil {
		logs.Error(err)
		return err
	}
	// 初始化数据库、校验器
	if err := transition.Init(); err != nil {
		logs.Error(err)
		return err
	}

	return nil
}

func main() {
	if err := AllInit(); err != nil {
		logs.Error(err)
		return
	}

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


	//http.HandleFunc("/server/changePasswordLink/send", controls.SendChangePasswordLink)
	//http.HandleFunc("/server/changePasswordLink/visit", controls.ChangePasswordLinkVisit)

	if err := http.ListenAndServe(fmt.Sprintf(":%d", env.Conf.Server.Port), nil); err != nil {
		logs.Error(err)
		return
	}
}
