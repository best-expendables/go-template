package handler

import (
	"net/http"

	_ "${REPO_HOST}/${PROJ_NAME}/model"
	"${REPO_HOST}/${PROJ_NAME}/service/${FILENAME}"
	"github.com/best-expendables/common-utils/util/request_parser"
	"github.com/best-expendables/common-utils/util/response"
)

type ${SERVICE_NAME}GetHandler struct {
	${CAMELIZED_NAME}Service ${FILENAME}.${SERVICE_NAME}Service
}

func New${SERVICE_NAME}GetHandler(
	${CAMELIZED_NAME}Service ${FILENAME}.${SERVICE_NAME}Service,
) ${SERVICE_NAME}GetHandler {
	return ${SERVICE_NAME}GetHandler{
		${CAMELIZED_NAME}Service: ${CAMELIZED_NAME}Service,
	}
}

// @Tags ${SERVICE_NAME}s
// @Router /v1/${PROJ_NAME}/${DASHED_NAME}s/{id} [GET]
// @ID ${SERVICE_NAME}GetHandler
// @Summary Get ${CAMELIZED_NAME}
// @Description Get ${CAMELIZED_NAME}
// @Param id path string true "${SERVICE_NAME} ID"
// @Success 200 {object} response.ApiResponse{data=model.${SERVICE_NAME}} "OK"
// @Failure 500 {object} response.Error "Internal Server Error"
func (h ${SERVICE_NAME}GetHandler) Handle(r *http.Request) response.ApiResponse {
	id := request_parser.URLParam(r, "id")
	output, err := h.${CAMELIZED_NAME}Service.Get(r.Context(), id)
	if err != nil {
		return response.ConvertServiceError(err)
	}
	return response.Ok(output)
}
