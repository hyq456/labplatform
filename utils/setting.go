package utils

import (
	"fmt"
	"gopkg.in/ini.v1"
)

var (
	AppMode  string
	HttpPort string
	Jwt      string

	Db         string
	DbHost     string
	DbPort     string
	DbUser     string
	DbPassWord string
	DbName     string

	AccessKey   string
	SecretKey   string
	Bucket      string
	QiniuServer string

	Mailaddress  string
	DialerAdress string
	DialerPort   string
	DialerUser   string
	DialerPW     string

	RedisHost     string
	RedisPort     string
	RedisPassWord string
	RedisDB       int
)

func init() {
	file, err := ini.Load("config/config.ini")
	if err != nil {
		fmt.Println("配置文件读写错误，请检查", err)
	}
	LoadSever(file)
	LoadData(file)
	LoadEmail(file)
	LoadRedis(file)
}

func LoadSever(file *ini.File) {
	AppMode = file.Section("server").Key("AppMode").MustString("debug")
	HttpPort = file.Section("server").Key("HttpPort").MustString("9999")
	Jwt = file.Section("sever").Key("Jwt").MustString("axcvbn")
}
func LoadData(file *ini.File) {
	DbHost = file.Section("database").Key("DbHost").MustString("localhost")
	DbPort = file.Section("database").Key("DbPort").MustString("3306")
	DbUser = file.Section("database").Key("DbUser").MustString("ginblog")
	DbPassWord = file.Section("database").Key("DbPassWord").String()
	DbName = file.Section("database").Key("DbName").MustString("ginblog")
}

func LoadEmail(file *ini.File) {
	Mailaddress = file.Section("email").Key("eaddress").String()
	DialerAdress = file.Section("email").Key("DialerAdress").String()
	DialerPort = file.Section("email").Key("DialerPort").String()
	DialerUser = file.Section("email").Key("DialerUser").String()
	DialerPW = file.Section("email").Key("DialerPW").String()
}

func LoadRedis(file *ini.File) {

	RedisHost = file.Section("redis").Key("RedisHost").String()
	RedisPort = file.Section("redis").Key("RedisPort").String()
	RedisPassWord = file.Section("redis").Key("RedisPassWord").String()
	RedisDB, _ = file.Section("redis").Key("RedisDB").Int()
}
