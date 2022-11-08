package gowalk

import (
	"errors"
	"github.com/cyj19/gowalk/config"
)

// Component 组件接口
type Component interface {
	Run() error   // 组件初始化
	Name() string // 获取组件名称
}

// 存放加载的组件实例
var components = make(map[string]Component)

// AddComponents 添加组件
func AddComponents(args ...Component) {
	for _, c := range args {
		components[c.Name()] = c
	}
}

// AddAndLoadComponents 添加并加载指定组件
func AddAndLoadComponents(args ...Component) error {
	// 全部组件初始化完成再添加到components
	for _, c := range args {
		// 加载组件配置
		err := config.GetConfig(c.Name(), c)
		if err != nil {
			return err
		}
		if err = c.Run(); err != nil {
			return err
		}
	}
	AddComponents(args...)

	return nil
}

// LoadAllComponents 加载全部组件
func LoadAllComponents() error {
	for _, c := range components {
		// 加载组件配置
		err := config.GetConfig(c.Name(), c)
		if err != nil {
			return err
		}
		if err = c.Run(); err != nil {
			return err
		}
	}

	return nil
}

// GetComponent 获取指定组件实例
func GetComponent(key string) (Component, error) {
	if c, ok := components[key]; ok {
		return c, nil
	} else {
		return nil, errors.New("Component:" + key + " does not find")
	}
}
