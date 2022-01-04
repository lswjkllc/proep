package src

import (
	"encoding/json"
	"sync"

	coms "github.com/lswjkllc/proep/src/commons"
)

// Container 相关
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

var (
	once      sync.Once
	container *Container
)

func GetContainer(path string) *Container {
	once.Do(func() {
		// 获取配置信息
		config := coms.GetConfig(path)
		// 初始化 Container
		container = &Container{BaseConfig: *config}
	})
	return container
}
