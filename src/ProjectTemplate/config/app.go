package config

import (
	"encoding/json"
	"github.com/kelseyhightower/envconfig"
	"time"
)

type AppConfig struct {
	RunMode          string `envconfig:"RUN_MODE" required:"true"`
	HttpPort         int    `envconfig:"HTTP_PORT" required:"true"`
	DisableAccessLog bool   `envconfig:"DISABLE_ACCESS_LOG" required:"false"`

	NewRelicAppName string `envconfig:"NEWRELIC_APPNAME" required:"true"`
	NewRelicLicense string `envconfig:"NEWRELIC_LICENSE" required:"true"`

	DBConfig
}

type DBConfig struct {
	DBMasterConfig
	Slaves                      DBSlaveConfigs `envconfig:"POSTGRES_SLAVES"`
	DbLogEnable                 bool           `envconfig:"POSTGRES_LOG_MODE_ENABLE"`
	DbMaxIdleConnection         int            `envconfig:"POSTGRES_MAX_IDLE_CONNECTION" required:"true"`
	DbMaxConnection             int            `envconfig:"POSTGRES_MAX_CONNECTION" required:"true"`
	DbMaxConnectionLifetime     time.Duration  `envconfig:"POSTGRES_MAX_CONNECTION_LIFETIME" required:"true"`
	DbMaxIdleConnectionLifetime time.Duration  `envconfig:"POSTGRES_MAX_IDLE_CONNECTION_LIFETIME" required:"true"`
}

type DBMasterConfig struct {
	DbHost string `envconfig:"POSTGRES_HOST" required:"true"`
	DbPort int    `envconfig:"POSTGRES_PORT" required:"true"`
	DbName string `envconfig:"POSTGRES_DB" required:"true"`
	DbUser string `envconfig:"POSTGRES_USER" required:"true"`
	DbPass string `envconfig:"POSTGRES_PASS" required:"true"`
}

type DBSlaveConfig struct {
	DbHost string `json:"host"`
	DbPort int    `json:"port"`
	DbName string `json:"db"`
	DbUser string `json:"user"`
	DbPass string `json:"pass"`
}

type DBSlaveConfigs []DBSlaveConfig

func (s *DBSlaveConfigs) Decode(value string) error {
	var slaves []DBSlaveConfig
	if err := json.Unmarshal([]byte(value), &slaves); err != nil {
		return err
	}
	*s = slaves
	return nil
}

func GetAppConfigFromEnv() AppConfig {
	var conf AppConfig
	envconfig.MustProcess("", &conf)
	return conf
}
