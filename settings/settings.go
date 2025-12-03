package settings

import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

// Conf 定义全局配置文件结构体
var Conf = new(AppConfig)

// AppConfig 全局配置结构体
type AppConfig struct {
	Name          string `mapstructure:"name"`
	Mode          string `mapstructure:"mode"`
	Port          int    `mapstructure:"port"`
	Version       string `mapstructure:"version"`
	StartTime     string `mapstructure:"start_time"`
	MachineID     int64  `mapstructure:"machine_id"`
	*AuthConfig   `mapstructure:"auth"`
	*LoggerConfig `mapstructure:"log"`
	*MySQLConfig  `mapstructure:"mysql"`
	*RedisConfig  `mapstructure:"redis"`
}

// AuthConfig 认证相关配置结构体
type AuthConfig struct {
	JwtExpire int64 `mapstructure:"jwt_expire"`
}

// LoggerConfig 日志配置结构体
type LoggerConfig struct {
	Level      string `mapstructure:"level"`
	FileName   string `mapstructure:"filename"`
	MaxSize    int    `mapstructure:"max_size"`
	MaxBackups int    `mapstructure:"max_backups"`
	MaxAge     int    `mapstructure:"max_age"`
}

// MySQLConfig mysql配置相关结构体
type MySQLConfig struct {
	Host         string `mapstructure:"host"`
	Port         int    `mapstructure:"port"`
	User         string `mapstructure:"user"`
	Password     string `mapstructure:"password"`
	DBName       string `mapstructure:"dbname"`
	MaxOpenConns int    `mapstructure:"max_open_conns"`
	MaxIdleConns int    `mapstructure:"max_idle_conns"`
}

// RedisConfig redis配置相关结构体
type RedisConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Password string `mapstructure:"password"`
	DB       int    `mapstructure:"db"`
	PoolSize int    `mapstructure:"pool_size"`
}

// Init 配置文件初始化
func Init() (err error) {
	// 设置viper读取路径
	// 相对于启动时项目的路径而言，该项目main.go跟配置文件在同一路径。
	viper.SetConfigFile("./config.yaml")
	// 读取配置文件
	if err = viper.ReadInConfig(); err != nil {
		zap.L().Error("viper.ReadInConfig failed, err:%v\n", zap.Error(err))
		return
	}
	// 绑定到结构体
	if err = viper.Unmarshal(Conf); err != nil {
		zap.L().Error("viper.Unmarshal failed, err:%v\n", zap.Error(err))
		panic(err)
	}
	// 监控配置文件变化
	viper.WatchConfig()
	// 配置文件变化则输出提示
	viper.OnConfigChange(func(e fsnotify.Event) {
		zap.L().Debug("config file changed")
		// 重新绑定
		if err = viper.Unmarshal(Conf); err != nil {
			zap.L().Error("viper.Unmarshal failed, err:%v\n", zap.Error(err))
			panic(err)
		}
	})
	return
}
