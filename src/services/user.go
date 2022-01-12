package services

import (
	"time"

	coms "github.com/lswjkllc/proep/src/commons"
	ms "github.com/lswjkllc/proep/src/models"
	"gorm.io/gorm"
)

type UserService struct {
	Config *coms.ConfigInfo
	Db     *gorm.DB
}

func (ucase UserService) whereById(id int) *gorm.DB {
	return ucase.Db.Where("id = ? and is_deleted = 0", id)
}

func (ucase UserService) CreateUser(user *ms.User) error {
	// 查找
	tx := ucase.Db.Create(user)

	return tx.Error
}

func (ucase UserService) GetUserById(id int) (ms.User, error) {
	user := ms.User{}
	err := ucase.whereById(id).First(&user).Error

	return user, err
}

func (ucase UserService) Save(user *ms.User) {
	ucase.Db.Save(user)
}

func (ucase UserService) DeleteUserById(id int, hard bool) error {
	if !hard {
		// 软珊: 快速
		value := time.Now().UnixNano() / 1e6
		return ucase.whereById(id).Updates(ms.User{IsDeleted: value}).Error
	}
	// 硬删: 耗时
	return ucase.whereById(id).Delete(ms.User{}).Error
}

func NewService(config *coms.ConfigInfo, db *gorm.DB) *UserService {
	return &UserService{Config: config, Db: db}
}
