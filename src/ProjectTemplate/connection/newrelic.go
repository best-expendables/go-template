package connection

import (
	"github.com/best-expendables/logger"
	"github.com/newrelic/go-agent"
)

type NewRelicConfig struct {
	AppName string
	License string
}

func CreateNewRelicApp(conf NewRelicConfig) newrelic.Application {
	newRelicConf := newrelic.NewConfig(
		conf.AppName,
		conf.License,
	)
	newRelicApp, err := newrelic.NewApplication(newRelicConf)
	if err != nil {
		logger.Error(err)
	}
	return newRelicApp
}
