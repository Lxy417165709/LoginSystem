package models

import "2_models/table"


type MainDb interface {
	Insert(table table.ITable) error
	Delete(table table.ITable, queryStr string, parameters ...interface{}) error
	Update(table table.ITable, queryStr string, parameters ...interface{}) error
	Select(table table.ITable, queryStr string, parameters ...interface{}) ([]table.ITable, error)
}
