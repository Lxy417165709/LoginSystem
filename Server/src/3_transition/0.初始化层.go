package transition

import (
	dtc "2_models/dataCenter"
	"2_models/pstsql"
	"2_models/pud"
	"2_models/rds"
	"2_models/vc"
	"github.com/astaxie/beego/logs"
)

// 数据库
var redis = &rds.Redis{}
var pgsql  = &pstsql.Pgsql{}
var dataCenter  = &dtc.DataCenter{}

// 校验器
var registerVrcManager *vc.RegisterEmailVrcChecker            // 注册验证码校验器
//var changePasswordVrcManager *vc.ChangePasswordEmailVrcChecker // 修改密码验证码校验器

// 图片上传与查看工具
var photoUploader *pud.PhotoUploader

// 初始化
func Init() error {
	if err := redis.Init();err!=nil{
		logs.Error(err)
		return err
	}

	if err := pgsql.Init(); err != nil {
		logs.Error(err)
		return err
	}

	dataCenter = dtc.NewDataCenter(redis,pgsql)
	registerVrcManager = vc.NewRegisterEmailVrcChecker(redis)
	//changePasswordVrcManager = vc.NewChangePasswordEmailVrcChecker(redis)
	return nil
}

// 关闭
func Close() error {
	if err := pgsql.Close(); err != nil {
		logs.Error(err)
		return err
	}
	if err := redis.Close(); err != nil {
		logs.Error(err)
		return err
	}
	return nil
}


