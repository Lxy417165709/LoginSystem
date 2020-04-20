package pstsql

import (
	"0_common/commonInterface"
	"github.com/astaxie/beego/logs"
)


func (p *Pgsql) Insert(tb commonInterface.ITable) error {
	logs.Info(getInsertSql(tb))
	sentence, values, err := getInsertSql(tb)
	if err != nil {
		return err
	}
	_, err = p.db.Exec(sentence, values...)
	return err
}

func (p *Pgsql) Delete(tb commonInterface.ITable,queryStr string, parameters ...interface{}) error {
	logs.Info(getDeleteSql(tb,queryStr, parameters...))
	sentence, values,err := getDeleteSql(tb,queryStr, parameters...)
	if err != nil {
		return err
	}
	_, err = p.db.Exec(sentence, values...)
	return err
}

func (p *Pgsql) Update(tb commonInterface.ITable,queryStr string, parameters ...interface{}) error {
	logs.Info(getUpdateSql(tb,queryStr, parameters...))
	sentence, values,err := getUpdateSql(tb,queryStr, parameters...)

	if err != nil {
		return err
	}
	_, err = p.db.Exec(sentence, values...)
	return err
}

func (p *Pgsql) Select(tb commonInterface.ITable, queryStr string, parameters ...interface{}) (result []commonInterface.ITable, err error) {
	logs.Info(getSelectSql(tb,queryStr, parameters...))
	sentence, values, err := getSelectSql(tb,queryStr, parameters...)

	if err != nil {
		return nil, err
	}

	rows, err := p.db.Query(sentence, values...)

	if rows == nil{
		return result, nil
	}
	for rows.Next() {
		if err = rows.Scan(getFieldAddr(tb)...); err != nil {
			return nil, err
		}
		result = append(result, tb)
	}
	return result, nil
}
