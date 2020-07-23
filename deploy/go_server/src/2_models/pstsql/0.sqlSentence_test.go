package pstsql

import (
	"2_models/table"
	"fmt"
	"testing"
)

func TestSql(t *testing.T) {
	uai := table.NewDefaultUai(1,"123","hello")
	upi := table.NewDefaultUpi(1,"456")
	fmt.Println(GetInsertSql(uai))
	fmt.Println(GetDeleteSql(uai,"where UserId=$1",1))
	fmt.Println(GetUpdateSql(uai,"where UserId=$1",1))
	fmt.Println(GetSelectSql(uai,"where UserId=$1",1))
	fmt.Println("---------------")
	fmt.Println(GetInsertSql(upi))
	fmt.Println(GetDeleteSql(upi,"where UserId=$1",2))
	fmt.Println(GetUpdateSql(upi,"where UserId=$1",2))
	fmt.Println(GetSelectSql(upi,"where UserId=$1",2))
}
