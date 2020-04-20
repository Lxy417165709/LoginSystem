package pstsql

import (
	"1_env"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)



type Pgsql struct {
	db *sql.DB
}

func (p *Pgsql) Init() error {
	var db *sql.DB
	// 开始连接数据库
	var err error
	if db, err = sql.Open(env.Conf.Postgresql.DriverName, fmt.Sprintf(
		`host=%s port=%d user=%s password=%s dbname=%s sslmode=%s`,
		env.Conf.Postgresql.Host,
		env.Conf.Postgresql.Port,
		env.Conf.Postgresql.User,
		env.Conf.Postgresql.Password,
		env.Conf.Postgresql.Dbname,
		env.Conf.Postgresql.Sslmode,
	)); err != nil {
		return err
	}

	// 设置数据库连接
	db.SetMaxIdleConns(env.Conf.Postgresql.MaxIdleConns)
	db.SetMaxOpenConns(env.Conf.Postgresql.MaxOpenConns)

	// ping 数据库
	if err := db.Ping(); err != nil {
		return err
	}
	p.db = db
	return nil
}

func (p *Pgsql) Close() error {
	return p.db.Close()
}
