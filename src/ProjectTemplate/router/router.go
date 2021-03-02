package router

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/best-expendables/logger"
	"github.com/go-chi/chi"
	httpSwagger "github.com/swaggo/http-swagger"
)

type HandlerFuncs struct {
	${SERVICE_NAME}GetHandler   http.HandlerFunc
	${SERVICE_NAME}PostHandle   http.HandlerFunc
	${SERVICE_NAME}PutHandle    http.HandlerFunc
	${SERVICE_NAME}PatchHandle  http.HandlerFunc
	${SERVICE_NAME}DeleteHandle http.HandlerFunc
}

func Create(handlerFuncs HandlerFuncs) http.Handler {
	api := chi.NewRouter()

	api.Group(func(r chi.Router) {
		r.Route("/v1/${FILENAME}s", func(r chi.Router) {
			r.Route("/", func(r chi.Router) {
				r.Get("/", handlerFuncs.${SERVICE_NAME}GetHandler)
				r.Post("/", handlerFuncs.${SERVICE_NAME}PostHandle)
				r.Put("/", handlerFuncs.${SERVICE_NAME}PutHandle)
				r.Patch("/", handlerFuncs.${SERVICE_NAME}PatchHandle)
				r.Delete("/", handlerFuncs.${SERVICE_NAME}DeleteHandle)
			})
		})
	})
	return api
}

func SwaggerHandler(rw http.ResponseWriter, _ *http.Request) {
	content, err := ioutil.ReadFile("build/swagger.json")
	if err != nil {
		logger.Error(err)
		rw.WriteHeader(500)
		return
	}

	baseURL := os.Getenv("API_BASE_URL")
	if baseURL != "" {
		host := strings.Split(baseURL, "//")[1]
		content = bytes.Replace(content, []byte("${REPO_HOST}/${PROJ_NAME}.local"), []byte(host), 1)
	}

	length, err := rw.Write(content)
	if err != nil {
		logger.Error(err)
		rw.WriteHeader(500)
		return
	}
	rw.Header().Add("Content-Type", "application/json")
	rw.Header().Add("Content-Length", string(length))
	rw.Header().Add("Cache-Control", "no-cache")
}

func SwaggerUIHandler(w http.ResponseWriter, r *http.Request) {
	handler := httpSwagger.Handler(
		httpSwagger.URL("/swagger/api-docs"),
	)
	handler.ServeHTTP(w, r)
}
