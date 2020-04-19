package rds

import (
	"1_env"
	"testing"
)

func TestRedis_Del(t *testing.T) {
	r := Redis{}
	t.Log(env.LoadConf("../../conf.ini"))
	t.Log(r.Init())
	//t.Log(r.Set("uai:uid:5",models.NewDefaultUai(5,"417165709","123456789"),120))
	t.Log(r.Get("uai:uid:6"))
}
