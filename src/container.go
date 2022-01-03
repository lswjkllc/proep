package main

import (
	"encoding/json"

	coms "github.com/lswjkllc/proep/src/commons"
)

type Container struct {
	BaseConfig coms.ConfigInfo `yaml:"config" json:"config"`
}

func (container Container) String() string {
	out, err := json.Marshal(container)
	if err != nil {
		return err.Error()
	}
	return string(out)
}

func GetContainer() *Container {
	// 获取配置信息
	config := coms.GetConfig()
	// 初始化 Container
	container := &Container{BaseConfig: *config}
	return container
}
