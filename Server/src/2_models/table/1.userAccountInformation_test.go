package table

import (
	"fmt"
	"testing"
)

func TestUserAccountInformation_GetSelectSql(t *testing.T) {
	uai:=UserAccountInformation{UserId:1,UserEmail:"11",UserPassword:"123"}
	fmt.Println(uai.GetSelectSql("where userEmail=$1","123456"))
}
