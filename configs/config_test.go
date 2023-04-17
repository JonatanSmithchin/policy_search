package configs

import (
	"fmt"
	"github.com/spf13/viper"
	"testing"
)

func TestInitConfig(t *testing.T) {
	InitConfig(".")
	fmt.Println(viper.Get("db.DriverName"))
}
