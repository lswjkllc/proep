package services

import (
	coms "github.com/lswjkllc/proep/src/commons"
	ms "github.com/lswjkllc/proep/src/models"
	"gorm.io/gorm"
)

type UserService struct {
	Config *coms.ConfigInfo
	Db     *gorm.DB
}

func (userUsecase UserService) CreateUser(user *ms.User) error {
	// 查找
	tx := userUsecase.Db.Create(user)

	return tx.Error
}

func (ucase UserService) GetUserById(id int) (ms.User, error) {
	user := ms.User{}
	err := ucase.Db.Where("id", id).First(&user).Error

	return user, err
}

func (ucase UserService) UpdateUserById(id int, user *ms.User) error {
	return ucase.Db.Updates(&user).Where("id", id).Error
}

func (ucase UserService) GetUser(user *ms.User) error {
	// 查找
	tx := ucase.Db.Find(user)
	return tx.Error
}

func NewService(config *coms.ConfigInfo, db *gorm.DB) *UserService {
	return &UserService{Config: config, Db: db}
}
