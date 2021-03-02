package handler

import (
	"net/http"

	"${REPO_HOST}/${PROJ_NAME}/service/${FILENAME}"
	"github.com/best-expendables/common-utils/util/request_parser"
	"github.com/best-expendables/common-utils/util/response"
)

type ${SERVICE_NAME}DeleteHandler struct {
	${CAMELIZED_NAME}Service ${FILENAME}.${SERVICE_NAME}Service
}

func New${SERVICE_NAME}DeleteHandler(
	${CAMELIZED_NAME}Service ${FILENAME}.${SERVICE_NAME}Service,
) ${SERVICE_NAME}DeleteHandler {
	return ${SERVICE_NAME}DeleteHandler{
		${CAMELIZED_NAME}Service: ${CAMELIZED_NAME}Service,
	}
}

// @Tags ${SERVICE_NAME}s
// @Router /v1/${PROJ_NAME}/${DASHED_NAME}s/{id} [DELETE]
// @ID ${SERVICE_NAME}DeleteHandler
// @Summary Delete ${CAMELIZED_NAME}
// @Description Delete ${CAMELIZED_NAME}
// @Param id path string true "${SERVICE_NAME} ID"
// @Success 200 "OK"
// @Failure 400 {object} response.Error "Bad Request"
// @Failure 422 {object} response.Error "Unprocessable Entity"
// @Failure 500 {object} response.Error "Internal Server Error"
func (h ${SERVICE_NAME}DeleteHandler) Handle(r *http.Request) response.ApiResponse {
	id := request_parser.URLParam(r, "id")
	err := h.${CAMELIZED_NAME}Service.Delete(r.Context(), id)
	if err != nil {
		return response.ConvertServiceError(err)
	}
	return response.Ok(nil)
}
