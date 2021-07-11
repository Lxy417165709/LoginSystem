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

func NewPgsql() *Pgsql{
	return &Pgsql{&sql.DB{}}
}
func (p *Pgsql) MdbInit(host string,port int,user,password,dbname,sslmode string,maxIdleConns,maxOpenConns int) error {
	// 开始连接数据库
	var err error
	if p.db, err = sql.Open(env.Conf.Postgresql.DriverName, fmt.Sprintf(
		`host=%s port=%d user=%s password=%s dbname=%s sslmode=%s`,
		host,
		port,
		user,
		password,
		dbname,
		sslmode,
	)); err != nil {
		return err
	}

	// 设置数据库连接
	p.db.SetMaxIdleConns(maxIdleConns)
	p.db.SetMaxOpenConns(maxOpenConns)

	// ping 数据库
	if err := p.db.Ping(); err != nil {
		return err
	}
	return nil
}
func (p *Pgsql) MdbClose() error {
	return p.db.Close()
}
