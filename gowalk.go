package gowalk

import (
	"fmt"
	"github.com/cyj19/gowalk/config"
	"github.com/cyj19/gowalk/logk"
	"os"
	"path/filepath"
)

func init() {

	workDir, _ := os.Getwd()
	configName := "config.yml"

	// 加载配置文件
	configPath := filepath.Join(workDir, configName)

	err := config.LoadConfig(configPath)
	if err != nil {
		panic(fmt.Sprintf("LoadConfig error: %+v \n", err))
	}

	logCfg := logk.LogConfig{}
	err = config.GetConfig("log", &logCfg)
	if err != nil {
		panic(fmt.Sprintf("GetConfig error: %+v \n", err))
	}

	// 初始化日志
	err = logk.SetupLog(workDir, logCfg)
	if err != nil {
		panic(fmt.Sprintf("SetupLog error: %+v \n", err))
	}
}

func Run(args ...Component) error {
	return AddAndLoadComponents(args...)
}
