package conf

import (
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"github.com/spf13/viper"
	"time"
	"wusthelper-mp-gateway/library/cache/redis"
	"wusthelper-mp-gateway/library/database"
)

const (
	DevEnv  = "dev"
	ProdEnv = "prod"
)

var (
	Conf = &Config{}
)

type Config struct {
	MiniProgram MpConf
	Server      ServerConf
	Wusthelper  WusthelperConf
	Database    database.Config
	Redis       redis.Config
}

type ServerConf struct {
	Env            string
	Port           int
	Address        string
	BaseUrl        string
	TokenSecret    string
	PasswordKey    string
	TokenTimeout   time.Duration
	LogLocation    string
	VersionLogFile string
}

type WusthelperConf struct {
	Upstream     string
	Timeout      time.Duration
	Proxy        string
	TokenKey     string
	AdminBaseUrl string
}

type MpConf struct {
	QQ struct {
		AppID  string
		Secret string
	}
	Wechat struct {
		AppID  string
		Secret string
	}
}

func Init() (err error) {
	viper.AddConfigPath(".")
	viper.AddConfigPath("./conf")
	viper.AddConfigPath("./config")
	viper.AddConfigPath("/etc/wusthelper-mp")
	viper.AddConfigPath("$HOME/.wusthelper-mp")

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(Conf)
	if err != nil {
		return
	}

	if Conf.Server.Env == DevEnv {
		jsonByte, _ := jsoniter.Marshal(Conf)
		fmt.Println(string(jsonByte))
	}

	return
}
