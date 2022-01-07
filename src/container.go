package src

import (
	"encoding/json"
	"sync"

	coms "github.com/lswjkllc/proep/src/commons"
	ms "github.com/lswjkllc/proep/src/models"
	ss "github.com/lswjkllc/proep/src/services"
	"gorm.io/gorm"
)

// Container 相关
type Container struct {
	BaseConfig  *coms.ConfigInfo `yaml:"config" json:"config"`
	DB          *gorm.DB         `yaml:"db" json:"db"`
	UserUsecase *ss.UserService  `yaml:"userUsecase" json:"userUsecase"`
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

func GetContainer() *Container {
	config := coms.GetConfig()
	once.Do(func() {
		// 获取 mysql 连接
		db := ms.InitDB(config, false)
		// 获取 user 服务
		userUsecase := ss.NewService(config, db)
		// 初始化 Container
		container = &Container{BaseConfig: config, DB: db, UserUsecase: userUsecase}
	})
	return container
}
