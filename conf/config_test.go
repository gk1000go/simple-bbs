package conf

import (
	"fmt"
	"testing"
)

func TestInitialConfig(t *testing.T) {
	fConf := "config.default.toml"
	err := InitialConfig(fConf)
	if err == nil {
		fmt.Println(DefaultConfig)
		fmt.Println(len(DefaultConfig.BaseDb.Slaves))
	}
}
