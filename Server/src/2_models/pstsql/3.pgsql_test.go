package pstsql

import (
	"1_env"
	"fmt"
	"github.com/astaxie/beego/logs"
	_ "github.com/lib/pq"
	"testing"
)

func TestPgsql_InsertUai(t *testing.T) {
	logs.SetLogFuncCallDepth(4)
	logs.EnableFuncCallDepth(true)

	logs.Info(env.LoadConf("../../conf.ini"))
	p := &Pgsql{}
	logs.Info(p.Init())
	for i:=140;i<145;i++{
		logs.Info(p.GenerateNewUser(fmt.Sprintf("%d",i),fmt.Sprintf("%d",i)))
	}
}

func TestPgsql_GetUai(t *testing.T) {
	logs.SetLogFuncCallDepth(4)
	logs.EnableFuncCallDepth(true)

	logs.Info(env.LoadConf("../../conf.ini"))
	p := &Pgsql{}
	logs.Info(p.Init())

	logs.Info(p.GetUai("105"))
}

func TestPgsql_Init(t *testing.T) {
	logs.SetLogFuncCallDepth(3)
	logs.EnableFuncCallDepth(true)

	logs.Info(env.LoadConf("../../conf.ini"))
	p := &Pgsql{}
	logs.Info(p.Init())


	logs.Info(p.UpdateLastLoginTime(87))
	logs.Info(p.UpdateUserName(87,"123"))
	logs.Info(p.UpdatePhotoUrl(87,"121113"))
	logs.Info(p.UpdatePassword(87,"127877887871113"))
}
