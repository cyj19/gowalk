package gowalk

import (
	"flag"
	"github.com/cyj19/gowalk/core"
	"log"
	"os"
	"path/filepath"
)

func init() {

	// 获取程序模式 dev/pro
	flag.StringVar(&core.EnvMode, "mode", "dev", "Program Environment Mode")
	// 获取工作目录
	flag.StringVar(&core.WorkDir, "wd", "", "Work Dir")
	flag.Parse()

	if core.WorkDir == "" {
		core.WorkDir, _ = os.Getwd()
	}

	configName := "config." + core.EnvMode + ".yml"

	// 加载配置文件
	configPath := filepath.Join(core.WorkDir, configName)

	err := core.LoadConfig(configPath)
	if err != nil {
		log.Fatalf("LoadConfig error: %+v \n", err)
	}
}

func Run(args ...core.Component) {
	_ = core.AddAndLoadComponents(args...)
}
