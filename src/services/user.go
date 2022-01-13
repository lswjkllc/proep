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

func (ucase UserService) whereNotDeleted() *gorm.DB {
	return ucase.Db.Where("is_deleted", 0)
}

func (ucase UserService) whereById(id int) *gorm.DB {
	return ucase.whereNotDeleted().Where("id", id)
}

func (ucase UserService) where(searchData map[string]interface{}) *gorm.DB {
	tx := ucase.whereNotDeleted()

	name, ok := searchData["name"]
	if ok {
		tx = tx.Where("name LIKE ?", name.(string)+"%")
	}
	age, ok := searchData["age"]
	if ok {
		tx = tx.Where("age", uint(age.(float64)))
	}
	gender, ok := searchData["gender"]
	if ok {
		tx = tx.Where("gender", gender)
	}
	email, ok := searchData["email"]
	if ok {
		tx = tx.Where("email", email)
	}

	return tx
}

func (ucase UserService) FindUsers(searchData map[string]interface{}) ([]ms.User, int64) {
	var users []ms.User
	offset, limit := getPageInfo(searchData)

	result := ucase.where(searchData).Order("id desc").Offset(offset).Limit(limit).Find(&users)

	return users, result.RowsAffected
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
