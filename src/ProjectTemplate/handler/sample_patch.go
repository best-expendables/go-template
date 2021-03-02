package handler

import (
	"net/http"

	_ "${REPO_HOST}/${PROJ_NAME}/model"
	"${REPO_HOST}/${PROJ_NAME}/service/${FILENAME}"
	"${REPO_HOST}/${PROJ_NAME}/service/${FILENAME}/dto"
	"github.com/best-expendables/common-utils/util/request_parser"
	"github.com/best-expendables/common-utils/util/response"
)

type ${SERVICE_NAME}PatchHandler struct {
	${CAMELIZED_NAME}Service ${FILENAME}.${SERVICE_NAME}Service
}

func New${SERVICE_NAME}PatchHandler(
	${CAMELIZED_NAME}Service ${FILENAME}.${SERVICE_NAME}Service,
) ${SERVICE_NAME}PatchHandler {
	return ${SERVICE_NAME}PatchHandler{
		${CAMELIZED_NAME}Service: ${CAMELIZED_NAME}Service,
	}
}

// @Tags ${SERVICE_NAME}s
// @Router /v1/${PROJ_NAME}/${DASHED_NAME}s/{id} [PATCH]
// @ID ${SERVICE_NAME}PatchHandler
// @Summary Patch ${CAMELIZED_NAME}
// @Description Patch ${CAMELIZED_NAME}
// @Param id path string true "${SERVICE_NAME} ID"
// @Param content body dto.${SERVICE_NAME}PatchUpdateInput true "${SERVICE_NAME} patch payload"
// @Success 200 {object} response.ApiResponse{data=model.${SERVICE_NAME}} "OK"
// @Failure 400 {object} response.Error "Bad Request"
// @Failure 422 {object} response.Error "Unprocessable Entity"
// @Failure 500 {object} response.Error "Internal Server Error"
func (h ${SERVICE_NAME}PatchHandler) Handle(r *http.Request) response.ApiResponse {
	var payload dto.${SERVICE_NAME}PatchUpdateInput
	if err := request_parser.DecodePayload(r, &payload); err != nil {
		return response.BadRequest(err)
	}
	id := request_parser.URLParam(r, "id")
	output, err := h.${CAMELIZED_NAME}Service.Patch(r.Context(), id, payload)
	if err != nil {
		return response.ConvertServiceError(err)
	}
	return response.Ok(output)
}
