package config

import (
	"fmt"
	"github.com/pelletier/go-toml"
	"log"
	"os"
)

var GlobalConfig *Config

type ServerConfig struct {
	IP          string `toml:"IP"`
	PORT        int    `toml:"PORT"`
	EnableTLS   bool   `toml:"EnableTLS"`
	TLSCertPath string `toml:"TLSCertPath"`
	TLSKeyPath  string `toml:"TLSKeyPath"`
	AdminInfo   struct {
		Username string `toml:"username"`
		Password string `toml:"password"`
	} `toml:"AdminInfo"`
	ApiAuthToken string `toml:"ApiAuthToken"`
}

type JavbusSiteConfig struct {
	EnableProxy bool     `toml:"EnableProxy"`
	ProxyUrl    string   `toml:"ProxyUrl"`
	JavbusUrls  []string `toml:"JavbusUrls"`
	UserAgents  []string `toml:"UserAgents"`
	UserCookies []struct {
		Username string `toml:"username"`
		Cookies  string `toml:"cookies"`
	} `toml:"UserCookies"`
}

type Config struct {
	Server     ServerConfig     `toml:"Server"`
	JavbusSite JavbusSiteConfig `toml:"JavbusSite"`
}

func (config *Config) initial(configFile string) {
	configPath := configFile

	// 打开配置文件
	file, err := os.Open(configPath)
	if err != nil {
		log.Fatalf("无法打开配置文件: %v", err)
	}
	defer file.Close()

	// 解析配置文件
	decoder := toml.NewDecoder(file)
	if err := decoder.Decode(config); err != nil {
		log.Fatalf("解析配置文件时出错: %v", err)
	}

	// 打印解析结果
	fmt.Printf("%+v\n", config)
}

func defaultConfig() *Config {

	return &Config{}
}

func NewConfig(configFile string) *Config {
	config := defaultConfig()
	config.initial(configFile)
	return config
}

//func init() {
//	configFile := "config.toml"
//	GlobalConfig = NewConfig(configFile)
//}
