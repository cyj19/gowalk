### gowalk
![](https://img.shields.io/github/go-mod/go-version/cyj19/gowalk/master) ![](https://img.shields.io/badge/viper-v1.13.0-blue) ![](https://img.shields.io/badge/zap-v1.17.0-blue) ![](https://img.shields.io/badge/lumberjack-v2.0.0-blue) ![](https://img.shields.io/github/license/cyj19/gowalk)


### 关于
gowalk 一个简单易用的Golang后端开发框架，开箱即用，通过加载组件获取数据库等功能。
> 名字来源 ：gowalk 意为去散步，希望gowalk可以像散步一样足够自由与灵活。



### 组件
> 开发中，持续更新



### 使用
1. gowalk设置了两个命令行参数`wd ` (工作目录，默认当前目录) 和 `mode` (环境模式：dev/prod， 默认dev模式)。

2. 配置文件名称约定为config.模式.yml，默认加载wd下的配置文件
3. 框架提供了[Component接口](https://github.com/cyj19/gowalk/blob/main/component.go)，[Logger接口](https://github.com/cyj19/gowalk/blob/main/logk/logger.go)，[Server接口](https://github.com/cyj19/gowalk/blob/main/transport/transport.go)，开发者可以自定义这些功能。
4. 示例  
```
type greeter struct {
    name string
}

// 实现component接口
func (g *greeter) Run() error {
    logk.Infof("hello, %s", g.name)
    return nil
}

func (g *greeter) Name() string {
    return "greeter"
}

var _ gowalk.Component = (*greeter)(nil)

func main() {
    // 加载组件
    _ = gowalk.Run(&greeter{name: "gowalk"})
    
    g, err := gowalk.GetComponent("greeter")
    if err != nil {
        logk.Fatal(err)
    }
    logk.Infof("greeter: %#v", g)
}

```
[示例源码](https://github.com/cyj19/gowalk/tree/main/examples/greeter)

### MIT License

    Copyright (c) 2022 cyj19