package commons

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"sync"

	"gopkg.in/yaml.v2"

	us "github.com/lswjkllc/proep/src/utils"
)

//解析yml文件
type ConfigInfo struct {
	CommonBase CommonBaseEntity `yaml:"commonbase" json:"commonbase"`
	LogBase    LogBaseEntity    `yaml:"logbase" json:"logbase"`
	DataBase   DataBaseEntity   `yaml:"database" json:"database"`
}

// 公共基础信息
type CommonBaseEntity struct {
	Name  string `yaml:"name" json:"name"`
	Addr  string `yaml:"addr" json:"addr"`
	Host  string `yaml:"host" json:"host"`
	Port  int    `yaml:"port" json:"port"`
	Area  string `yaml:"area" json:"area"`
	Env   string `yaml:"env" json:"env"`
	Debug bool   `yaml:"debug" json:"debug"`
}

// 日志配置信息
type LogBaseEntity struct {
	Level      string `yaml:"level" json:"level"`
	Path       string `yaml:"path" json:"path"`
	MaxSize    int    `yaml:"maxsize" json:"maxsize"`
	MaxBackups int    `yaml:"maxbackups" json:"maxbackups"`
	MaxAge     int    `yaml:"maxage" json:"maxage"`
	Compress   bool   `yaml:"compress" json:"compress"`
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

func (myde MysqlDataEntity) GetDsn() string {
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

func (info *ConfigInfo) ReadFile(path string) {
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

func GetConfigByPath(path string) *ConfigInfo {
	config := &ConfigInfo{}
	config.ReadFile(path)
	config.CommonBase.Addr = us.JoinStrings(config.CommonBase.Host, ":", strconv.Itoa(config.CommonBase.Port))
	return config
}

// 初始化 config
var (
	once   sync.Once
	config *ConfigInfo
)

func GetConfig() *ConfigInfo {
	once.Do(func() {
		config = GetConfigByPath("./config/config.yaml")
	})
	return config
}
