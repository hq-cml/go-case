package mysql

import "testing"

func TestInsert(t *testing.T) {
	//var val []interface{} = []interface{}{"2", "hx"}
	vals := []interface{}{"2", "hx"}
	err := Insert("root:123456@tcp(127.0.0.1:3306)", "test", "INSERT INTO test VALUES( ?, ? )", vals)
	if err != nil {
		t.Fatal("Insert error:", err.Error())
	}
}

func TestSelect(t *testing.T) {
	err := Select("root:123456@tcp(127.0.0.1:3306)", "test", "select * from test")
	if err != nil {
		t.Fatal("select error:", err.Error())
	}
}
