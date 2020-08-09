package conf

import (
	"fmt"
	"testing"
)

func TestInitSqlContent(t *testing.T) {
	fConf := "config.sql.toml"

	err := InitSqlContent(fConf)
	if err == nil {
		fmt.Printf("%+v\n", SqlAll)
		fmt.Println(len(SqlAll.SqlText))
	}
}
