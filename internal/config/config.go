package config

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"sync"
)

// AppConfig 结构体，用于保存程序的所有配置信息
type AppConfig struct {
	Name         string `mapstructure:"name"`
	Mode         string `mapstructure:"mode"`
	Version      string `mapstructure:"version"`
	Port         int    `mapstructure:"port"`
	StartTime    string `mapstructure:"start_time"`
	MachineID    int64  `mapstructure:"machine_id"`
	*LogConfig   `mapstructure:"log"`
	*MySQLConfig `mapstructure:"mysql"`
	*RedisConfig `mapstructure:"redis"`
}

type LogConfig struct {
	Level      string `mapstructure:"level"`
	Filename   string `mapstructure:"filename"`
	MaxSize    int    `mapstructure:"max_size"`
	MaxAge     int    `mapstructure:"max_age"`
	MaxBackups int    `mapstructure:"max_backups"`
}

type MySQLConfig struct {
	Host         string `mapstructure:"host"`
	User         string `mapstructure:"user"`
	Password     string `mapstructure:"password"`
	DbName       string `mapstructure:"db_name"`
	Port         int    `mapstructure:"port"`
	MaxOpenConns int    `mapstructure:"max_open_conns"`
	MaxIdleConns int    `mapstructure:"max_idle_conns"`
}

type RedisConfig struct {
	Host     string `mapstructure:"host"`
	Password string `mapstructure:"password"`
	Port     int    `mapstructure:"port"`
	DB       int    `mapstructure:"db"`
	PoolSize int    `mapstructure:"pool_size"`
}

var (
	configOnce sync.Once
	conf       *AppConfig
)

func GetConfig() *AppConfig {

	// 使用 sync.Once 保证配置只加载一次
	configOnce.Do(func() {
		conf = new(AppConfig) // 这里直接修改全局变量 conf
		configFileName := "./configs/config.yaml"
		// 设置配置文件路径和文件名
		viper.SetConfigFile(configFileName)
		// 读取配置文件
		err := viper.ReadInConfig()
		if err != nil {
			// 读取配置文件失败
			fmt.Printf("viper.ReadInConfig() failed, err:%v\n", err)
			return
		}

		// 将配置文件内容反序列化到 conf 对象中
		if err := viper.Unmarshal(conf); err != nil {
			fmt.Printf("viper.Unmarshal failed, err:%v\n", err)
			return
		}

		// 配置文件变化监听
		viper.WatchConfig()
		viper.OnConfigChange(func(in fsnotify.Event) {
			fmt.Println("配置文件修改了...")
			if err := viper.Unmarshal(conf); err != nil {
				fmt.Printf("viper.Unmarshal failed, err:%v\n", err)
			}
		})
	})

	// 返回配置对象
	return conf
}
