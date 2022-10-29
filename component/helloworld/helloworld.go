package helloworld

import "log"

type Instance struct {
	Message string
}

var componentName = "helloworld"

func (i *Instance) Run() error {
	log.Println(componentName + " Component starts to initialize...")
	return nil
}

func (i *Instance) Name() string {
	return componentName
}

func (i *Instance) GetSettings() interface{} {
	return i
}
