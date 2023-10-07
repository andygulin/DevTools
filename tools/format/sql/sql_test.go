package sql

import "testing"

func TestSQLFormat_Format(t *testing.T) {
	obj := new(SQLFormat)
	input := "select * from test_table where name like '%a%' order by id desc limit 10"
	output, err := obj.Format(input)
	if err != nil {
		return
	}
	t.Log(output)
}

func TestSQLFormat_FormatFile(t *testing.T) {

}
