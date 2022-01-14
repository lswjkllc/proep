package models

import "encoding/json"

type Goods struct {
	Id        uint   `gorm:"primaryKey;autoIncrement;comment:主键" json:"id" form:"id" query:"id"`
	CreatedAt int64  `gorm:"autoCreateTime:milli;not null;comment:创建时间（毫秒级）" json:"createdAt" form:"createdAt" query:"createdAt"`
	UpdatedAt int64  `gorm:"autoUpdateTime:milli;not null;comment:更新时间（毫秒级）" json:"updateAt" form:"updateAt" query:"updateAt"`
	Name      string `gorm:"size:50;not null;uniqueIndex:unq_idx_name,priority:1;index:idx_name;comment:名字" json:"name" form:"name" query:"name"`
	IsDeleted int64  `gorm:"not null;uniqueIndex:unq_idx_name,priority:2;comment:是否删除" json:"isDeleted" form:"isDeleted" query:"isDeleted"`
	Count     int    `gorm:"type:int(11);not null;comment:数量" json:"count" form:"count" query:"count"`
}

func (Goods) TableName() string {
	return "goods"
}

func (goods Goods) String() string {
	out, err := json.Marshal(goods)
	if err != nil {
		return err.Error()
	}
	return string(out)
}
