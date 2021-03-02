package handler

import (
	"github.com/best-expendables/common-utils/util/response"
	"gorm.io/gorm"
	"net/http"
)

type HealthCheckGet struct {
	db *gorm.DB
}

func NewHealthCheckGet(db *gorm.DB) HealthCheckGet {
	return HealthCheckGet{
		db: db,
	}
}

func (h HealthCheckGet) Handle(r *http.Request) response.ApiResponse {
	status := "Error"
	statusCode := http.StatusInternalServerError
	rawDB, err := h.db.DB()
	if err == nil && rawDB.Ping() == nil {
		status = "OK"
		statusCode = http.StatusOK
	}
	return response.ApiResponse{
		Code: statusCode,
		Data: map[string]string{"status": status},
	}
}
