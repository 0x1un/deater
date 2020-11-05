package conf

import (
	"github.com/spf13/viper"
	"log"
	"strings"
)

func ReadAgentConfig(filepath string, searchPath ...string) *viper.Viper {
	if strings.HasSuffix(filepath, "yaml") || strings.HasSuffix(filepath, "yml") {
		viper.SetConfigFile(filepath)
	}
	// 读取多个搜索路径
	for _, path := range searchPath {
		if path != "" {
			viper.AddConfigPath(path)
		}
	}
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(err)
	}
	return viper.GetViper()
}
