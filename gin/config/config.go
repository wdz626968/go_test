package config

import (
	"go_test/gin/global"
	"sync/atomic"
)

var (
	appConfig   atomic.Value // *Config
	dbConfig    atomic.Value // *DBConfig
	cacheConfig atomic.Value // *CacheConfig
	jwtConfig   atomic.Value // *JWTConfig
)

type Config struct {
	Name string `mapstructure:"name"`
	Port int    `mapstructure:"port"`
}

type DBConfig struct {
	Host            string `mapstructure:"host"`
	Port            int    `mapstructure:"port"`
	User            string `mapstructure:"user"`
	Password        string `mapstructure:"password"`
	Name            string `mapstructure:"name"`
	MaxIdleConns    int    `mapstructure:"max_idle_conns"`
	MaxOpenConns    int    `mapstructure:"max_open_conns"`
	ConnMaxLifetime int    `mapstructure:"conn_max_lifetime"`
}

type CacheConfig struct {
	ArticleExpire int `mapstructure:"article_expire"`
	LikeExpire    int `mapstructure:"like_expire"`
}

type JWTConfig struct {
	Secret      string `mapstructure:"secret"`
	ExpireHours int    `mapstructure:"expire_hours"`
}

// GetAppConfig 原子读取应用配置
func GetAppConfig() *Config {
	if config := appConfig.Load(); config != nil {
		return config.(*Config)
	}
	return nil
}

// GetDBConfig 原子读取数据库配置
func GetDBConfig() *DBConfig {
	if config := dbConfig.Load(); config != nil {
		return config.(*DBConfig)
	}
	return nil
}

// GetCacheConfig 原子读取缓存配置
func GetCacheConfig() *CacheConfig {
	if config := cacheConfig.Load(); config != nil {
		return config.(*CacheConfig)
	}
	return nil
}

// GetJWTConfig 原子读取JWT配置
func GetJWTConfig() *JWTConfig {
	if config := jwtConfig.Load(); config != nil {
		return config.(*JWTConfig)
	}
	return nil
}

func InitConfig() {
	global.InitDB(InitDB())
}
