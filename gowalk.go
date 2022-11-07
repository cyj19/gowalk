package gowalk

import (
	"flag"
	"log"
	"os"
	"path/filepath"
	"sync"
)

var (
	EnvMode string // 模式 dev/pro
	WorkDir string // 工作目录
	once    sync.Once
)

func Init() {

	// 获取程序模式 dev/prod
	flag.StringVar(&EnvMode, "mode", "dev", "Program Environment Mode")
	// 获取工作目录
	flag.StringVar(&WorkDir, "wd", "", "Work Dir")
	flag.Parse()

	if WorkDir == "" {
		WorkDir, _ = os.Getwd()
	}

	configName := "config." + EnvMode + ".yml"

	// 加载配置文件
	configPath := filepath.Join(WorkDir, configName)

	err := LoadConfig(configPath)
	if err != nil {
		log.Fatalf("LoadConfig error: %+v \n", err)
	}

	// 初始化日志
	SetUp()
}

func Run(args ...Component) {
	once.Do(func() {
		Init()
	})
	_ = AddAndLoadComponents(args...)
}
