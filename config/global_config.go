package config

import (
	"flag"
	"github.com/spf13/viper"
)

var (
	configPath   string
	GlobalConfig Config
)

type Config struct {
	Server Server `yaml:"server"`
	MySql  MySql  `yaml:"mysql"`
	Redis  Redis  `yaml:"redis"`
}

type Server struct {
	Port        int    `yaml:"port"`
	EnablePprof bool   `yaml:"enable_pprof"`
	LogLevel    string `yaml:"log_level"`
}

type MySql struct {
	Host        string `yaml:"host"`
	Port        int    `yaml:"port"`
	User        string `yaml:"user"`
	Password    string `yaml:"password"`
	Database    string `yaml:"database"`
	Charset     string `yaml:"charset"`
	ShowSql     bool   `yaml:"show_sql"`
	MaxOpenConn int    `yaml:"max_open_conn"`
	MaxIdleConn int    `yaml:"max_idle_conn"`
}

type Redis struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Password string `yaml:"password"`
	Db       int    `yaml:"db"`
	MaxIdle  int    `yaml:"max_idle"`
	MaxOpen  int    `yaml:"max_open"`
}

func init() {
	flag.StringVar(&configPath, "c", "application.yml", "default config path")
}

func InitConfig() Config {
	var (
		err        error
		tempConfig = &Config{}
		vipConfig  = viper.New()
	)
	flag.Parse()
	tempConfig, err = getConfig(vipConfig)
	if err != nil {
		panic(err)
	}
	return *tempConfig
}

func getConfig(vipConfig *viper.Viper) (*Config, error) {
	vipConfig.SetConfigFile(configPath)
	if err := vipConfig.ReadInConfig(); err != nil {
		return nil, err
	}
	if err := vipConfig.Unmarshal(&GlobalConfig); err != nil {
		return nil, err
	}
	return &GlobalConfig, nil
}
