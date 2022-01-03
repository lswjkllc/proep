package commons

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

//解析yml文件
type ConfigInfo struct {
	CommBase CommonBaseEntity `yaml:"commbase"`
	DataBase DataBaseEntity   `yaml:"database"`
}

// 公共基础信息
type CommonBaseEntity struct {
	Name string `yaml:"name"`
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

// 数据库基础信息
type DataBaseEntity struct {
	MysqlData MysqlDataEntity `yaml:"mysql"`
	RedisData RedisDataEntity `yaml:"redis"`
}

// mysql 数据库信息
type MysqlDataEntity struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Name     string `yaml:"name"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}

// redis 数据库信息
type RedisDataEntity struct {
	Host    string `yaml:"host"`
	Port    int    `yaml:"port"`
	Name    string `yaml:"name"`
	Timeout int    `yaml:"timeout"`
}

func (info *ConfigInfo) Init() {
	// 读取文件
	yamlFile, err := ioutil.ReadFile("./config/config.yaml")
	if err != nil {
		panic(err.Error())
	}
	// 解析文件
	err = yaml.Unmarshal(yamlFile, info)
	if err != nil {
		panic(err.Error())
	}
}
