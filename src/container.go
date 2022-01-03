package main

import (
	coms "github.com/lswjkllc/proep/src/commons"
)

func GetConfig() *coms.ConfigInfo {
	config := &coms.ConfigInfo{}
	config.Init()
	return config
}
