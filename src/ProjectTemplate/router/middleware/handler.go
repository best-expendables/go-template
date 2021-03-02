package middleware

import (
	"github.com/best-expendables/common-utils/util/response"
	"net/http"
)

// Func Handler function
type HandlerFunc func(r *http.Request) response.ApiResponse

// MakeHandler Create handler
func MakeHandler(handlerFunc HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		res := handlerFunc(r)
		response.RenderJson(w, res)
	}
}
