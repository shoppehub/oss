package conf

import (
	_ "embed"
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

//go:embed .env
var defaultEnvString string

const (
	// 应用的端口
	PORT = "PORT"
)

var config *viper.Viper

func init() {

	config = viper.New()

	config.SetConfigType("yaml")
	config.ReadConfig(strings.NewReader(defaultEnvString))

	config.SetConfigName("config")
	//viper.AddConfigPath("conf") // 还可以在工作目录中查找配置
	config.AddConfigPath(".") // 还可以在工作目录中查找配置

	err := config.MergeInConfig() // 查找并读取配置文件
	if err != nil {               // 处理读取配置文件的错误
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

}

func GetString(key string) string {
	return config.GetString(key)
}
