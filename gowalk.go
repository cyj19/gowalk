package gowalk

import (
	"flag"
	"fmt"
	"github.com/cyj19/gowalk/config"
	"github.com/cyj19/gowalk/logx"
	"os"
	"path/filepath"
	"sync"
)

var (
	envMode string // 模式 dev/prod
	workDir string // 工作目录
	once    sync.Once
)

func initConfig() {

	// 获取程序模式 dev/prod
	flag.StringVar(&envMode, "mode", "dev", "Program Environment Mode")
	// 获取工作目录
	flag.StringVar(&workDir, "wd", "", "Work Dir")
	flag.Parse()

	if workDir == "" {
		workDir, _ = os.Getwd()
	}

	configName := "config." + envMode + ".yml"

	// 加载配置文件
	configPath := filepath.Join(workDir, configName)

	err := config.LoadConfig(configPath)
	if err != nil {
		panic(fmt.Sprintf("LoadConfig error: %+v \n", err))
	}

	logCfg := logx.LogConfig{}
	err = config.GetConfig("log", &logCfg)
	if err != nil {
		panic(fmt.Sprintf("GetConfig error: %+v \n", err))
	}

	// 初始化日志
	err = logx.SetupLog(workDir, logCfg)
	if err != nil {
		panic(fmt.Sprintf("SetupLog error: %+v \n", err))
	}
}

func Run(args ...Component) error {
	once.Do(func() {
		initConfig()
	})
	return AddAndLoadComponents(args...)
}
