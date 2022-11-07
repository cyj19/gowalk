package gowalk

import "github.com/spf13/viper"

var vp *viper.Viper

// LoadConfig 加载配置文件到viper
func LoadConfig(path string) error {
	// 每次加载框架配置文件，都新建viper实例
	vp = viper.New()
	vp.SetConfigFile(path)
	return vp.ReadInConfig()
}

// GetConfig 根据配置文件的key获取配置内容
func GetConfig(key string, v interface{}) error {
	return vp.UnmarshalKey(key, v)
}
