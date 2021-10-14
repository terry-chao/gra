package viper

import (
	"fmt"
	"github.com/spf13/viper"
)

// 定义config结构体
type HostConfig struct {
	Host Host
}

// json中的嵌套对应结构体的嵌套
type Host struct {
	Address string
	Port    int
}

func getHostConfig() *HostConfig {
	config := viper.New()
	config.AddConfigPath("./viper")
	config.SetConfigName("host_config")
	config.SetConfigType("json")
	if err := config.ReadInConfig(); err != nil {
		panic(err)
	}
	fmt.Println(config.GetString("host.address"))
	fmt.Println(config.GetString("host.port"))

	//直接反序列化为Struct
	var configjson HostConfig
	if err := config.Unmarshal(&configjson); err != nil {
		fmt.Println(err)
	}

	return &configjson
}

type MysqlConfig struct {
	MysqlInfo MysqlInfo
}

type MysqlInfo struct {
	Username   string
	Password   string
	Address    string
	Table_name string
}

func GetMysqlConfig() *MysqlConfig {
	config := viper.New()
	config.AddConfigPath("./viper")
	config.SetConfigName("mysql_config")
	config.SetConfigType("json")
	if err := config.ReadInConfig(); err != nil {
		panic(err)
	}

	//直接反序列化为Struct
	var configjson MysqlConfig
	if err := config.Unmarshal(&configjson); err != nil {
		fmt.Println(err)
	}

	return &configjson
}
