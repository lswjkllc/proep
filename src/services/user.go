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
	conn, err := userUsecase.Db.DB()
	if err != nil {
		return err
	}
	defer conn.Close()

	// 创建
	userUsecase.Db.Create(user)

	return err
}

func NewService(config *coms.ConfigInfo, db *gorm.DB) *UserService {
	return &UserService{Config: config, Db: db}
}
