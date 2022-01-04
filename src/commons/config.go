package commons

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

//解析yml文件
type ConfigInfo struct {
	CommonBase CommonBaseEntity `yaml:"commonbase" json:"commonbase"`
	DataBase   DataBaseEntity   `yaml:"database" json:"database"`
}

// 公共基础信息
type CommonBaseEntity struct {
	Name string `yaml:"name" json:"name"`
	Host string `yaml:"host" json:"host"`
	Port int    `yaml:"port" json:"port"`
	Env  string `yaml:"env" json:"env"`
}

// 数据库基础信息
type DataBaseEntity struct {
	MysqlData MysqlDataEntity `yaml:"mysql" json:"mysql"`
	RedisData RedisDataEntity `yaml:"redis" json:"redis"`
}

// mysql 数据库信息
type MysqlDataEntity struct {
	Host     string `yaml:"host" json:"host"`
	Port     int    `yaml:"port" json:"port"`
	Name     string `yaml:"name" json:"name"`
	User     string `yaml:"user" json:"user"`
	Password string `yaml:"password" json:"password"`
}

func (myde MysqlDataEntity) GetDns() string {
	dns := fmt.Sprintf(
		"%s:%s@(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		myde.User, myde.Password, myde.Host, myde.Port, myde.Name)
	return dns
}

// redis 数据库信息
type RedisDataEntity struct {
	Host    string `yaml:"host" json:"host"`
	Port    int    `yaml:"port" json:"port"`
	Name    string `yaml:"name" json:"name"`
	Timeout int    `yaml:"timeout" json:"timeout"`
}

func (info *ConfigInfo) Init(path string) {
	// 读取文件
	yamlFile, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err.Error())
	}
	// 解析文件
	err = yaml.Unmarshal(yamlFile, info)
	if err != nil {
		panic(err.Error())
	}
}

func GetConfig(path string) *ConfigInfo {
	config := &ConfigInfo{}
	config.Init(path)
	return config
}
