package services

import (
	coms "github.com/lswjkllc/proep/src/commons"
	ms "github.com/lswjkllc/proep/src/models"
	ul "github.com/lswjkllc/proep/src/utils"
	"gorm.io/gorm"
)

type UserService struct {
	Config *coms.ConfigInfo
	Db     *gorm.DB
}

func (userUsecase UserService) CreateUser(data map[string]interface{}) (*ms.User, error) {
	conn, err := userUsecase.Db.DB()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	// 初始化 User
	user := &ms.User{}
	// map to struct
	err = ul.Map2struct(data, user)
	if err != nil {
		return nil, err
	}

	// 创建
	userUsecase.Db.Create(user)

	return user, err
}

func NewService(config *coms.ConfigInfo, db *gorm.DB) *UserService {
	return &UserService{Config: config, Db: db}
}
