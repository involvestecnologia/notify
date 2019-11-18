package config

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

var (
	configFile    = ".notify"
	configFileExt = "yml"
	defaultConfig = map[string]interface{}{
		"webhook": map[string]string{
			"mm":    "localhost/hooks/123",
			"slack": "localhost/hooks/123",
		},
	}
)

func Load() *viper.Viper {
	v := viper.New()
	for key, value := range defaultConfig {
		v.SetDefault(key, value)
	}
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	v.AutomaticEnv()
	v.SetConfigName(configFile)
	v.SetConfigType(configFileExt)
	if home := v.GetString("HOME"); home != "" {
		v.AddConfigPath(home)
	}
	v.AddConfigPath(".")
	if err := v.ReadInConfig(); err != nil {
		v.WriteConfigAs(fmt.Sprintf("%s.%s", configFile, configFileExt))
	}
	return v
}
