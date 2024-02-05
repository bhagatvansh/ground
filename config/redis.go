package config

import (
	"context"
	"fmt"
	"os"

	"github.com/techrail/ground/constants"
	"github.com/techrail/ground/constants/exitCode"
)

type redis struct {
	Main redisConfig
}

type redisConfig struct {
	Enabled                            bool
	Url                                string
	Username                           string
	Password                           string
	OperationMode                      string
	MaxActiveConnections               int
	MaxIdleConnections                 int
	IdleTimeoutInSeconds               int
	CrashAppOnConnectionFailure        bool
	ConnectRetryIntervalInSeconds      int
	AutoExpireTopLevelKeysAfterSeconds int
	AppNamespace                       string
	Address                            string
}

func init() {
	config.Redis = redis{
		Main: redisConfig{
			Enabled:                            constants.RedisEnabled,
			Url:                                constants.RedisDefaultUrl,
			Username:                           constants.RedisDefaultUsername,
			Password:                           constants.RedisDefaultPassword,
			OperationMode:                      constants.RedisDefaultOperationMode,
			MaxActiveConnections:               constants.RedisMaxActiveConnections,
			MaxIdleConnections:                 constants.RedisMaxIdleConnections,
			IdleTimeoutInSeconds:               constants.RedisIdleTimeoutInSeconds,
			CrashAppOnConnectionFailure:        constants.RedisCrashAppOnConnectionFailure,
			ConnectRetryIntervalInSeconds:      constants.RedisConnectRetryIntervalInSeconds,
			AutoExpireTopLevelKeysAfterSeconds: constants.RedisAutoExpireTopLevelKeysAfterSeconds,
			AppNamespace:                       constants.RedisAppNameSpace,
			Address:                            constants.RedisDefaultAddr,
		},
	}
}

var ctx = context.Background()

func initializeRedisConfig() {
	config.Redis.Main.Enabled = envOrViperOrDefaultBool("redis.main.enabled", config.Redis.Main.Enabled)
	config.Redis.Main.Url = envOrViperOrDefaultString("redis.main.url", config.Redis.Main.Url)
	config.Redis.Main.Username = envOrViperOrDefaultString("redis.main.username", config.Redis.Main.Username)
	config.Redis.Main.Password = envOrViperOrDefaultString("redis.main.password", config.Redis.Main.Password)
	config.Redis.Main.OperationMode = envOrViperOrDefaultString("redis.main.operationmode", config.Redis.Main.OperationMode)
	config.Redis.Main.MaxActiveConnections = int(envOrViperOrDefaultInt64("redis.main.maxActiveConnections", int64(config.Redis.Main.MaxActiveConnections)))
	config.Redis.Main.MaxIdleConnections = int(envOrViperOrDefaultInt64("redis.main.maxIdleConnections", int64(config.Redis.Main.MaxIdleConnections)))
	config.Redis.Main.IdleTimeoutInSeconds = int(envOrViperOrDefaultInt64("redis.main.idleTimeoutInSeconds", int64(config.Redis.Main.IdleTimeoutInSeconds)))
	config.Redis.Main.CrashAppOnConnectionFailure = envOrViperOrDefaultBool("redis.main.crashAppOnConnectionFailure", config.Redis.Main.CrashAppOnConnectionFailure)
	config.Redis.Main.ConnectRetryIntervalInSeconds = int(envOrViperOrDefaultInt64("redis.main."+
		"connectRetryIntervalInSeconds", int64(config.Redis.Main.ConnectRetryIntervalInSeconds)))
	config.Redis.Main.AutoExpireTopLevelKeysAfterSeconds = int(envOrViperOrDefaultInt64("redis.main.autoExpireTopLevelKeysAfterSeconds", int64(config.Redis.Main.AutoExpireTopLevelKeysAfterSeconds)))
	config.Redis.Main.AppNamespace = envOrViperOrDefaultString("redis.main.appNamespace", config.Redis.Main.AppNamespace)
	config.Redis.Main.Address = envOrViperOrDefaultString("redis.main.address", config.Redis.Main.Address)

	// TODO: work on it and implement it
	if config.Redis.Main.Enabled &&
		config.Redis.Main.OperationMode != "auto" &&
		config.Redis.Main.OperationMode != "cluster" &&
		config.Redis.Main.OperationMode != "standalone" {
		fmt.Printf("P#1MQUNR - Invalid redis operation mode. Cannot proceed.")
		os.Exit(exitCode.RedisConnectionFailed)
	}
}
