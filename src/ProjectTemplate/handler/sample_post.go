package handler

import (
	"net/http"

	_ "${REPO_HOST}/${PROJ_NAME}/model"
	"${REPO_HOST}/${PROJ_NAME}/service/${FILENAME}"
	"${REPO_HOST}/${PROJ_NAME}/service/${FILENAME}/dto"
	"github.com/best-expendables/common-utils/util/request_parser"
	"github.com/best-expendables/common-utils/util/response"
)

type ${SERVICE_NAME}PostHandler struct {
	${CAMELIZED_NAME}Service ${FILENAME}.${SERVICE_NAME}Service
}

func New${SERVICE_NAME}PostHandler(
	${CAMELIZED_NAME}Service ${FILENAME}.${SERVICE_NAME}Service,
) ${SERVICE_NAME}PostHandler {
	return ${SERVICE_NAME}PostHandler{
		${CAMELIZED_NAME}Service: ${CAMELIZED_NAME}Service,
	}
}

// @Tags ${SERVICE_NAME}s
// @Router /v1/${PROJ_NAME}/${DASHED_NAME}s [POST]
// @ID ${SERVICE_NAME}PostHandler
// @Summary Post ${CAMELIZED_NAME}
// @Description Post ${CAMELIZED_NAME}
// @Param content body dto.${SERVICE_NAME}CreateInput true "${SERVICE_NAME} post payload"
// @Success 200 {object} response.ApiResponse{data=model.${SERVICE_NAME}} "OK"
// @Failure 400 {object} response.Error "Bad Request"
// @Failure 422 {object} response.Error "Unprocessable Entity"
// @Failure 500 {object} response.Error "Internal Server Error"
func (h ${SERVICE_NAME}PostHandler) Handle(r *http.Request) response.ApiResponse {
	var payload dto.${SERVICE_NAME}CreateInput
	if err := request_parser.DecodePayload(r, &payload); err != nil {
		return response.BadRequest(err)
	}
	output, err := h.${CAMELIZED_NAME}Service.Create(r.Context(), payload)
	if err != nil {
		return response.ConvertServiceError(err)
	}
	return response.Ok(output)
}
