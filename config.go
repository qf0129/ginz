package ginz

import (
	"encoding/json"
	"os"
	"reflect"

	"github.com/sirupsen/logrus"
)

type Configuration struct {
	DbEngine   string
	SqliteFile string

	DbHost     string
	DbPort     uint
	DbUser     string
	DbPsd      string
	DbDatabase string

	// error|warn|info, default=warn
	DbLogLevel string

	DefaultPageSize int
	DBPrimaryKey    string

	AppHost    string
	AppPort    uint
	AppMode    string
	AppTimeout uint

	// panic|fatal|error|warn|info|debug|trace
	LogLevel string

	// 加密算法密钥
	Secret string

	TokenKey string
	// 令牌过期时间，单位秒
	TokenExpiredTime int64

	Custom map[string]any
}

func (c *Configuration) Set(k string, v any) {
	c.Custom[k] = v
}

func (c *Configuration) Get(k string) any {
	return c.Custom[k]
}

func (c *Configuration) GetStr(k string) (r string) {
	if val := c.Custom[k]; val != nil {
		r, _ = val.(string)
	}
	return
}

func (c *Configuration) GetInt(k string) (r int) {
	if val := c.Custom[k]; val != nil {
		r, _ = val.(int)
	}
	return
}

func (c *Configuration) GetInt64(k string) (r int64) {
	if val := c.Custom[k]; val != nil {
		r, _ = val.(int64)
	}
	return
}

func (c *Configuration) GetUint(k string) (r uint) {
	if val := c.Custom[k]; val != nil {
		r, _ = val.(uint)
	}
	return
}

func (c *Configuration) GetBool(k string) (r bool) {
	if val := c.Custom[k]; val != nil {
		r, _ = val.(bool)
	}
	return
}

func (c *Configuration) GetFloat64(k string) (r float64) {
	if val := c.Custom[k]; val != nil {
		r, _ = val.(float64)
	}
	return
}

func (c *Configuration) Remove(k string) {
	delete(c.Custom, k)
}

var Config = Configuration{
	DbEngine:   "sqlite",
	SqliteFile: "db.sqlite",

	DbHost:     "127.0.0.1",
	DbPort:     3306,
	DbUser:     "root",
	DbPsd:      "root",
	DbDatabase: "test",
	DbLogLevel: "warn",

	DBPrimaryKey:    "id",
	DefaultPageSize: 10,

	AppHost:    "",
	AppPort:    8080,
	AppMode:    "debug",
	AppTimeout: 60,
	LogLevel:   "debug",

	Secret:           "Abcd@123",
	TokenKey:         "tk",
	TokenExpiredTime: 7200,

	Custom: make(map[string]any),
}

func LoadConfig() {
	// 读取json文件
	data, err := os.ReadFile("config.json")
	if err != nil {
		logrus.Warn("Can not find config.json")
	}

	// 解析内置配置项
	err = json.Unmarshal(data, &Config)
	if err != nil {
		logrus.Warn("json unmarshal failed, err:", err)
	}

	// 解析自定义配置项
	var m map[string]any
	err = json.Unmarshal(data, &m)
	if err != nil {
		logrus.Warn("json unmarshal failed, err:", err)
	}

	// 遍历自定义配置项，不是内置项则放到Custom里
	for k, v := range m {
		_, ok := reflect.TypeOf(Config).FieldByName(k)
		if !ok {
			Config.Set(k, v)
		}
	}
}
