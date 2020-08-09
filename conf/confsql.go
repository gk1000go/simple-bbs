package conf

import (
	"fmt"
	"github.com/BurntSushi/toml"
)

type SqlContent struct {
	SqlText map[string]string `toml:"SqlText"`
}

var SqlAll SqlContent

func InitSqlContent(fSql string) error {

	if _, err := toml.DecodeFile(fSql, &SqlAll); err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}