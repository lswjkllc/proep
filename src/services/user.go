package services

import (
	coms "github.com/lswjkllc/proep/src/commons"
	"gorm.io/gorm"
)

type UserService struct {
	Config *coms.ConfigInfo
	DB     *gorm.DB
}
