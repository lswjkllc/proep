package src

import (
	"encoding/json"
	"fmt"
	"sync"

	"github.com/go-redis/redis/v8"
	coms "github.com/lswjkllc/proep/src/commons"
	ms "github.com/lswjkllc/proep/src/models"
	ss "github.com/lswjkllc/proep/src/services"
	"gorm.io/gorm"
)

// Container 相关
type Container struct {
	BaseConfig   *coms.ConfigInfo `yaml:"config" json:"config"`
	DB           *gorm.DB         `yaml:"db" json:"db"`
	Cache        *redis.Client    `yaml:"cache" json:"cache"`
	UserUsecase  *ss.UserService  `yaml:"userUsecase" json:"userUsecase"`
	GoodsUsecase *ss.GoodsService `yaml:"goodsUsecase" json:"goodsUsecase"`
}

func (container Container) String() string {
	out, err := json.Marshal(container)
	if err != nil {
		return err.Error()
	}
	return string(out)
}

func (container *Container) Close() {
	fmt.Println("清理 Container ...")
	// 清理 redis 连接
	container.Cache.Close()
	// 清理 mysql 连接
}

var (
	once      sync.Once
	container *Container
)

func GetContainer() *Container {
	config := coms.GetConfig()
	once.Do(func() {
		// 获取 mysql 连接
		db := ms.InitDB(&config.DataBase.MysqlData, config.CommonBase.Debug)
		// 获取 redis 连接
		cache := ms.InitCache(&config.DataBase.RedisData)
		// 获取 user 服务
		userUsecase := ss.NewUserService(config, db)
		// 获取 goods 服务
		goodsUsecase := ss.NewGoodsService(config, db, cache)
		// 初始化 Container
		container = &Container{
			BaseConfig: config, DB: db, Cache: cache,
			UserUsecase: userUsecase, GoodsUsecase: goodsUsecase}
	})
	return container
}
