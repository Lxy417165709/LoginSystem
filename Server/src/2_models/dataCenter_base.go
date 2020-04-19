package models

import (
	"2_models/table"
	"encoding/json"
)

type DataCenter struct {
	cache  Cache
	mainDb MainDb
}

func NewDataCenter(cache Cache, mainDb MainDb) *DataCenter {
	return &DataCenter{
		cache,
		mainDb,
	}
}

func (dbc DataCenter) Insert(key string, table table.ITable, expiredTime int) error {
	if err := dbc.mainDb.Insert(table); err != nil {
		return err
	}
	bytes, err := json.Marshal(&table)
	if err != nil {
		return err
	}
	return dbc.cache.Set(key, bytes, expiredTime)
}
func (dbc DataCenter) Delete(key string, table table.ITable, queryStr string, parameters ...interface{}) error {
	if err := dbc.mainDb.Delete(table, queryStr, parameters...); err != nil {
		return err
	}
	return dbc.cache.Del(key)
}
func (dbc DataCenter) Update(key string, table table.ITable, queryStr string, parameters ...interface{}) error {
	if err := dbc.mainDb.Update(table, queryStr, parameters...); err != nil {
		return err
	}
	return dbc.cache.Del(key)
}
func (dbc DataCenter) Select(key string, table table.ITable, queryStr string, parameters ...interface{}) (results []table.ITable, err error) {

	resultBytes, err := dbc.cache.Get(key)
	if err != nil {
		return nil, err
	}
	// 缓存有结果时
	if len(resultBytes) != 0 {
		if err := json.Unmarshal(resultBytes, &table); err != nil {
			return nil, err
		}
		results = append(results, table)
		return results, nil
	}
	// 缓存没结果时、错误时
	if results, err = dbc.mainDb.Select(table, queryStr, parameters...); err != nil || len(results) == 0 {
		return nil, err
	}

	bytes, err := json.Marshal(&results[0])
	if err != nil {
		return nil, err
	}

	if err := dbc.cache.Set(key, bytes, 120); err != nil {
		return nil, err
	}
	return results, nil
}
