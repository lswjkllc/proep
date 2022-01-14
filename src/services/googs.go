package services

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
	coms "github.com/lswjkllc/proep/src/commons"
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

func (gcase GoodsService) FlashSale(count int, key string) error {
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
		if v >= count {
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
	// watch（预扣）
	err := gcase.Cache.Watch(ctx, increment, key)

	return err
}

func NewGoodsService(config *coms.ConfigInfo, db *gorm.DB, cache *redis.Client) *GoodsService {
	return &GoodsService{Config: config, Db: db, Cache: cache}
}
