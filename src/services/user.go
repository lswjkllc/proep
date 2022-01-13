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

type UserList struct {
	Users  []ms.User `json:"users" yaml:"users"`
	Total  int64     `json:"total" yaml:"total"`
	Offset int       `json:"offset" yaml:"Offset"`
	Limit  int       `json:"limit" yaml:"limit"`
}

func (ucase UserService) whereNotDeleted() *gorm.DB {
	return ucase.Db.Where("is_deleted", 0)
}

func (ucase UserService) whereById(id int) *gorm.DB {
	return ucase.whereNotDeleted().Where("id", id)
}

func (ucase UserService) where(searchData map[string]interface{}) *gorm.DB {
	// 查询 未删除
	tx := ucase.whereNotDeleted()
	// 模糊查询 姓名
	name, ok := searchData["name"]
	if ok {
		tx = tx.Where("name LIKE ?", name.(string)+"%")
	}
	// 查询 年龄
	age, ok := searchData["age"]
	if ok {
		tx = tx.Where("age", uint(age.(float64)))
	}
	// 查询 性别
	gender, ok := searchData["gender"]
	if ok {
		tx = tx.Where("gender", gender)
	}
	// 查询 邮箱
	email, ok := searchData["email"]
	if ok {
		tx = tx.Where("email", email)
	}

	return tx
}

func (ucase UserService) FindUsers(searchData map[string]interface{}) UserList {
	// 声明
	var users []ms.User
	var count int64
	// 获取分页信息
	offset, limit := getPageInfo(searchData)
	// 查找
	ucase.where(searchData).Order("id desc").Offset(offset).Limit(limit).Find(&users).Count(&count)

	return UserList{Users: users, Total: count, Offset: offset, Limit: limit}
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
