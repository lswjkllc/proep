package models

import (
	"encoding/json"

	_ "gorm.io/driver/mysql"
)

// 定义一个数据模型(user表)
// 列名是字段名的蛇形小写(PassWd->pass_word)
type User struct {
	Id        uint   `gorm:"primaryKey;autoIncrement;comment:主键" json:"id"`
	CreatedAt int64  `gorm:"autoCreateTime:milli;not null;comment:创建时间（毫秒级）" json:"createdAt"`
	UpdatedAt int64  `gorm:"autoUpdateTime:milli;not null;comment:更新时间（毫秒级）" json:"updateAt"`
	Name      string `gorm:"size:50;not null;index:idx_name;comment:名字" json:"name"`
	Gender    string `gorm:"size:3;not null;comment:性别" json:"gender"`
	Age       uint   `gorm:"size:1;not null;comment:年龄" json:"age"`
	Birthday  string `gorm:"size:20;not null;comment:生日" json:"birthday"`
	Email     string `gorm:"type:varchar(50);not null;uniqueIndex:unq_idx_email;comment:邮箱（唯一）" json:"email"`
	PassWord  string `gorm:"type:varchar(25);not null;comment:密码" json:"password"`
}

func (user User) String() string {
	out, err := json.Marshal(user)
	if err != nil {
		return err.Error()
	}
	return string(out)
}
