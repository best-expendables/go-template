package handler

import (
	"net/http"

	"${REPO_HOST}/${PROJ_NAME}/service/${FILENAME}"
	"${REPO_HOST}/${PROJ_NAME}/service/${FILENAME}/dto"
	"github.com/best-expendables/common-utils/util/request_parser"
	"github.com/best-expendables/common-utils/util/response"
)

type ${SERVICE_NAME}ListHandler struct {
	${CAMELIZED_NAME}Service ${FILENAME}.${SERVICE_NAME}Service
}

func New${SERVICE_NAME}ListHandler(
	${CAMELIZED_NAME}Service ${FILENAME}.${SERVICE_NAME}Service,
) ${SERVICE_NAME}ListHandler {
	return ${SERVICE_NAME}ListHandler{
		${CAMELIZED_NAME}Service: ${CAMELIZED_NAME}Service,
	}
}

// @Tags ${SERVICE_NAME}s
// @Router /v1/${PROJ_NAME}/${DASHED_NAME}s [GET]
// @ID ${SERVICE_NAME}ListHandler
// @Summary List ${CAMELIZED_NAME}
// @Description List ${CAMELIZED_NAME}
// @Param orderBy query []string false "custom order by description"
// @Param request query dto.${SERVICE_NAME}GetListFilter true "${SERVICE_NAME} get list filter"
// @Success 200 {object} response.ApiResponse{data=[]model.${SERVICE_NAME}} "OK"
// @Failure 400 {object} response.Error "Bad Request"
// @Failure 500 {object} response.Error "Internal Server Error"
func (h ${SERVICE_NAME}ListHandler) Handle(r *http.Request) response.ApiResponse {
	f := dto.New${SERVICE_NAME}GetListFilter()
	if err := request_parser.DecodeURLParam(r, f); err != nil {
		return response.BadRequest(err)
	}
	output, err := h.${CAMELIZED_NAME}Service.List(r.Context(), f)
	if err != nil {
		return response.ConvertServiceError(err)
	}
	return response.Ok(output)
}
