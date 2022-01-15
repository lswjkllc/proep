package services

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
	coms "github.com/lswjkllc/proep/src/commons"
	ms "github.com/lswjkllc/proep/src/models"
	"gorm.io/gorm"
)

type GoodsService struct {
	Config *coms.ConfigInfo `json:"config" yaml:"config"`
	Db     *gorm.DB         `json:"db" yaml:"db"`
	Cache  *redis.Client    `json:"cache" yaml:"cache"`
}

type SaleError struct {
	Value string
}

func (se SaleError) Error() string {
	return se.Value
}

func (gcase GoodsService) cacheIncrToMax(key string, maxcount int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// increment
	increment := func(tx *redis.Tx) error {
		// 先查询下当前 watch 监听的 key 的值 v
		v, err := tx.Get(ctx, key).Int()
		if err != nil && err != redis.Nil {
			return err
		}
		// 当 缓存值 v >= count 设定值, 表示 秒杀结束
		if v >= maxcount {
			return &SaleError{"秒杀结束"}
		}
		// 如果 key的值 v 没有改变的话, TxPipelined 函数才会调用成功
		_, err = tx.TxPipelined(ctx, func(pipe redis.Pipeliner) error {
			// v++
			pipe.Incr(ctx, key)
			return nil
		})
		return err
	}
	// watch
	return gcase.Cache.Watch(ctx, increment, key)
}

func (gcase GoodsService) whereNotDeleted() *gorm.DB {
	return gcase.Db.Model(&ms.Goods{}).Where("is_deleted", 0)
}

func (gcase GoodsService) whereByName(name string) *gorm.DB {
	return gcase.whereNotDeleted().Where("name", name)
}

func (gcase GoodsService) dbDescCountByName(name string) error {
	return gcase.whereByName(name).Update("count", gorm.Expr("count - ?", 1)).Error
}

func (gcase GoodsService) FlashSale(name string, key string, maxcount int) error {
	// 预扣
	err := gcase.cacheIncrToMax(key, maxcount)
	if err != nil {
		return err
	}
	// 数据库扣除
	err = gcase.dbDescCountByName(name)

	return err
}

func NewGoodsService(config *coms.ConfigInfo, db *gorm.DB, cache *redis.Client) *GoodsService {
	return &GoodsService{Config: config, Db: db, Cache: cache}
}
