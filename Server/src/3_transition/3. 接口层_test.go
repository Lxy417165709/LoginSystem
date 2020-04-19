package transition

import (
	"1_env"
	"github.com/astaxie/beego/logs"
	"testing"
)

func TestLoginCheck(t *testing.T) {
	t.Log(env.LoadConf("../conf.ini"))
	t.Log(InitDBs())
	t.Log(LoginCheck("417165709@qq.com","QIQINGCHANG666."))
}

func TestUploadPhotoCheck(t *testing.T) {
	t.Log(env.LoadConf("../conf.ini"))
	t.Log(InitDBs())

	//t.Log(UploadPhotoCheck(5,))

}

func TestSendLink(t *testing.T) {
	t.Log(env.LoadConf("../conf.ini"))
	t.Log(InitDBs())
	t.Log(InitVrcChecker())
	logs.Info(sendLink("417165709@qq.com",60))
}
