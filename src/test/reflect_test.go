package test

import (
	"dao"
	"fmt"
	"reflect"
	"testing"
)

// TestTet 测试
func TestTet(t *testing.T) {
	var up []dao.UserProfile
	err := dao.Query(&up, "select id, acc_name, chinese_name, scholar_field, introduction from user_profile where id >= ? and id <= ?", 1, 5)
	rows, err1 := dao.Exec("update user_profile set chinese_name = ? where id = ?", "方文崇 updata", 1)
	fmt.Printf("%+v %+v\n", rows, err1)
	if err != nil {
		fmt.Printf("%+v\n", up)
	}
	// tr := reflect.ValueOf(up)
	// tet(tr)
}

func tet(tt reflect.Value) {
	fmt.Printf("%+v\n", tt.Type())
	newSlice := reflect.MakeSlice(tt.Type(), 100, 100)
	fmt.Printf("%+v\n", newSlice)
}
