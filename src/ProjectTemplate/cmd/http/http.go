package main

import (
	"errors"
	"fmt"
	"net/http"

	"${REPO_HOST}/${PROJ_NAME}/config"
	"${REPO_HOST}/${PROJ_NAME}/connection"
	"${REPO_HOST}/${PROJ_NAME}/handler"
	"${REPO_HOST}/${PROJ_NAME}/repository/postgresql"
	internalRouter "${REPO_HOST}/${PROJ_NAME}/router"
	internalMiddleware "${REPO_HOST}/${PROJ_NAME}/router/middleware"
	"${REPO_HOST}/${PROJ_NAME}/service/${FILENAME}"
	"${REPO_HOST}/${PROJ_NAME}/util/validation"
	"github.com/best-expendables/common-utils/util/response"
	utilValidator "github.com/best-expendables/common-utils/util/validation"
	"github.com/best-expendables/logger"
	"github.com/best-expendables/router"
	"github.com/go-chi/chi"
	chiMiddleware "github.com/go-chi/chi/middleware"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// @title Swagger ${PROJ_NAME} Service
// @version 1.0
func main() {
	appConf := config.GetAppConfigFromEnv()

	// Logger factory
	var loggerFactory logger.Factory
	if appConf.RunMode == "dev" {
		loggerFactory = logger.NewLoggerFactory(logger.DebugLevel)
	} else {
		loggerFactory = logger.NewLoggerFactory(logger.InfoLevel)
	}
	postgresConnection := connection.NewPostgresConnection(appConf.DBConfig)
	db := postgresConnection.CreateDB()

	// init NewRelic Application
	newRelicConf := connection.NewRelicConfig{
		AppName: appConf.NewRelicAppName,
		License: appConf.NewRelicLicense,
	}
	newRelicApp := connection.CreateNewRelicApp(newRelicConf)

	// Validation
	validator := validation.NewValidator()
	internalValidator := utilValidator.NewInternalValidator(validator)

	// Router
	routerConfig := router.Configuration{
		LoggerFactory: loggerFactory,
		NewrelicApp:   newRelicApp,
	}
	routerConfig.AccessLog.Disable = appConf.DisableAccessLog
	r, err := router.New(routerConfig)
	if err != nil {
		logger.Emergency(err)
		panic(err)
	}

	${CAMELIZED_NAME}Repo := postgresql.New${SERVICE_NAME}Repository(db)

	${CAMELIZED_NAME}Service := ${FILENAME}.New${SERVICE_NAME}Service(${CAMELIZED_NAME}Repo, internalValidator)

	handlerFuncs := internalRouter.HandlerFuncs{
		${SERVICE_NAME}GetHandler:   internalMiddleware.MakeHandler(handler.New${SERVICE_NAME}GetHandler(${CAMELIZED_NAME}Service).Handle),
		${SERVICE_NAME}PostHandle:   internalMiddleware.MakeHandler(handler.New${SERVICE_NAME}PostHandler(${CAMELIZED_NAME}Service).Handle),
		${SERVICE_NAME}PutHandle:    internalMiddleware.MakeHandler(handler.New${SERVICE_NAME}PutHandler(${CAMELIZED_NAME}Service).Handle),
		${SERVICE_NAME}PatchHandle:  internalMiddleware.MakeHandler(handler.New${SERVICE_NAME}PatchHandler(${CAMELIZED_NAME}Service).Handle),
		${SERVICE_NAME}DeleteHandle: internalMiddleware.MakeHandler(handler.New${SERVICE_NAME}DeleteHandler(${CAMELIZED_NAME}Service).Handle),
	}

	r.Group(func(r chi.Router) {
		api := internalRouter.Create(handlerFuncs)
		r.Mount("/api", api)
		r.Get("/swagger/api-docs", internalRouter.SwaggerHandler)
		r.Get("/swagger/*", internalRouter.SwaggerUIHandler)
		r.Mount("/metrics", promhttp.Handler())
		r.Mount("/debug", chiMiddleware.Profiler())
		r.Get("/health-check", internalMiddleware.MakeHandler(handler.NewHealthCheckGet(db).Handle))
		r.NotFound(func(rw http.ResponseWriter, _ *http.Request) {
			res := response.ErrorResponse(errors.New("404 Not Found"), 404)
			response.RenderJson(rw, res)
		})
	})

	logger.Info("Start listening...")
	if err := http.ListenAndServe(fmt.Sprintf(":%v", appConf.HttpPort), r); err != nil {
		logger.Emergency(err)
		panic(err)
	}
}
