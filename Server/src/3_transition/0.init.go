package transition

import (
	"1_env"
	dtc "2_models/dataCenter"
	"2_models/pstsql"
	"2_models/pud"
	"2_models/rds"
	"2_models/vc"
	"github.com/astaxie/beego/logs"
)

type FPI struct {
	dataCenter         *dtc.DataCenter
	registerVrcManager *vc.RegisterEmailVrcChecker
	photoUploader      *pud.PhotoUploader
}

var fpi = &FPI{
	dtc.NewDataCenter(rds.NewRedis(), pstsql.NewPgsql()),
	vc.NewRegisterEmailVrcChecker(rds.NewRedis()),
	pud.NewPhotoUploader(rds.NewRedis()),
}

func GetFPI() *FPI {
	return fpi
}

func Init() error {
	if err := fpi.dataCenter.Init(
		env.Conf.Postgresql.Host,
		env.Conf.Postgresql.Port,
		env.Conf.Postgresql.User,
		env.Conf.Postgresql.Password,
		env.Conf.Postgresql.Dbname,
		env.Conf.Postgresql.Sslmode,
		env.Conf.Postgresql.MaxIdleConns,
		env.Conf.Postgresql.MaxOpenConns,
		env.Conf.Redis.Network,
		env.Conf.Redis.Host,
		env.Conf.Redis.Port,
	); err != nil {
		panic(err)
	}
	if err := fpi.registerVrcManager.Init(
		env.Conf.EmailServer.User,
		env.Conf.EmailServer.Password,
		"qq",
		env.Conf.Redis.Network,
		env.Conf.Redis.Host,
		env.Conf.Redis.Port,
	); err != nil {
		panic(err)
	}
	if err := fpi.photoUploader.Init(
		env.Conf.Redis.Network,
		env.Conf.Redis.Host,
		env.Conf.Redis.Port,
	); err != nil {
		panic(err)
	}
	return nil
}

func Close() error {
	if err := fpi.dataCenter.Close(); err != nil {
		logs.Error(err)
		return err
	}
	if err := fpi.registerVrcManager.Close(); err != nil {
		logs.Error(err)
		return err
	}
	if err := fpi.photoUploader.Close(); err != nil {
		logs.Error(err)
		return err
	}
	return nil
}
