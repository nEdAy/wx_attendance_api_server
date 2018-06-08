package config

import (
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/naoina/toml"
)

var Conf *Config

// Config 系统根配置
type Config struct {
	Mode      string       `toml:"mode"`
	Http      *HTTPConfig  `toml:"http"`
	MySQL     *MySQLConfig `toml:"mysql"`
}

// HTTPConfig http监听配置
type HTTPConfig struct {
	Address string `toml:"address"`
	Port    int    `toml:"port"`
}

// Init http配置
func (hc *HTTPConfig) Init() {
	if hc.Address == "" {
		hc.Address = "0.0.0.0"
	}
	if hc.Port == 0 {
		hc.Port = 443
	}
}

// NewConfig 初始化一个配置文件对象
func NewConfig(configPath string) (*Config, error) {
	if configPath == "" {
		configPath = fmt.Sprintf(".%sdata%sconfig.toml", string(os.PathSeparator), string(os.PathSeparator))
	}
	log.Println(configPath)

	f, err := os.Open(configPath)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	config := new(Config)
	if err := toml.NewDecoder(f).Decode(config); err != nil {
		return nil, err
	}

	if config.Http == nil {
		return nil, errors.New("http配置不能为空")
	}
	config.Http.Init()
	Conf = config
	return config, nil
}

// Duration 用于日志文件解析出时间段
type Duration struct {
	time.Duration
}

// UnmarshalText implements encoding.TextUnmarshaler
func (d *Duration) UnmarshalText(text []byte) error {
	var err error
	d.Duration, err = time.ParseDuration(string(text))
	return err
}
