package transition

import (
	"2_models"
	"2_models/pstsql"
	"2_models/pud"
	"2_models/rds"
	"2_models/vc"
)

// 数据库
var redis = &rds.Redis{}
var pgsql  = &pstsql.Pgsql{}
var dataCenter  = &models.DataCenter{}

// 校验器
var registerVrcManager *vc.RegisterEmailVrcChecker            // 注册验证码校验器
var changePasswordVrcManager *vc.ChangePasswordEmailVrcChecker // 修改密码验证码校验器

// 图片上传与查看工具
var photoUploader *pud.PhotoUploader

// 初始化
func Init() error {
	if err := redis.Init();err!=nil{
		return err
	}

	if err := pgsql.Init(); err != nil {
		return err
	}

	dataCenter = models.NewDataCenter(redis,pgsql)
	registerVrcManager = vc.NewRegisterEmailVrcChecker(redis)
	changePasswordVrcManager = vc.NewChangePasswordEmailVrcChecker(redis)
	return nil
}

// 关闭
func Close() error {
	if err := pgsql.Close(); err != nil {
		return err
	}
	return redis.Close()
}


