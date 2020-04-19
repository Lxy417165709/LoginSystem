package pstsql

import (
	"2_models/table"
	"database/sql"
	"fmt"
	"github.com/astaxie/beego/logs"
)

type Dao struct {
	db *sql.DB
}

func (ud *Dao) Insert(table table.ITable) error {

	sentence, values, err := table.GetInsertSql()
	if err != nil {
		return err
	}
	_, err = ud.db.Exec(sentence, values...)
	return err
}

func (ud *Dao) Delete(table table.ITable,queryStr string, parameters ...interface{}) error {
	sentence, err := table.GetDeleteSql(queryStr, parameters...)
	if err != nil {
		return err
	}
	_, err = ud.db.Exec(sentence)
	return err

}

func (ud *Dao) Update(table table.ITable,queryStr string, parameters ...interface{}) error {
	logs.Info(table)
	sentence, values, err := table.GetUpdateSql(queryStr, parameters...)

	if err != nil {
		return err
	}
	_, err = ud.db.Exec(sentence, values...)
	return err
}

func (ud *Dao) Select(table table.ITable, queryStr string, parameters ...interface{}) (result []table.ITable, err error) {
	if table==nil{
		return nil,fmt.Errorf("can't know the type Of table")
	}

	sentence, values, err := table.GetSelectSql(queryStr, parameters...)
	fmt.Println(sentence, values, err )
	if err != nil {
		return nil, err
	}

	rows, err := ud.db.Query(sentence, values...)

	if rows == nil{
		return result, nil
	}
	for rows.Next() {
		if err = rows.Scan(table.GetFieldAddr()...); err != nil {
			return nil, err
		}
		result = append(result, table)
	}
	return result, nil
}




