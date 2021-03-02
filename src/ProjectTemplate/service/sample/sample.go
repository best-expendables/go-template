package ${FILENAME}

import (
	"${REPO_HOST}/${PROJ_NAME}/model"
	"${REPO_HOST}/${PROJ_NAME}/repository"
	"${REPO_HOST}/${PROJ_NAME}/service/${FILENAME}/dto"
	"github.com/best-expendables/common-utils/util/validation"
	"context"
)

type ${SERVICE_NAME}Service interface {
	List(ctx context.Context, filter *dto.${SERVICE_NAME}GetListFilter) ([]model.${SERVICE_NAME}, error)
	Get(ctx context.Context, id string, preloadFields ...string) (*model.${SERVICE_NAME}, error)
	Create(ctx context.Context, input dto.${SERVICE_NAME}CreateInput) (*model.${SERVICE_NAME}, error)
	Put(ctx context.Context, id string, input dto.${SERVICE_NAME}PutUpdateInput) (*model.${SERVICE_NAME}, error)
	Patch(ctx context.Context, id string, input dto.${SERVICE_NAME}PatchUpdateInput) (*model.${SERVICE_NAME}, error)
	Delete(ctx context.Context, id string) error
}

type ${CAMELIZED_NAME}Service struct {
	${CAMELIZED_NAME}Repo repository.${SERVICE_NAME}Repository
	validator  validation.Validator
}

func New${SERVICE_NAME}Service(
	${CAMELIZED_NAME}Repo repository.${SERVICE_NAME}Repository,
	validator validation.Validator,
) ${SERVICE_NAME}Service {
	return ${CAMELIZED_NAME}Service{
		${CAMELIZED_NAME}Repo: ${CAMELIZED_NAME}Repo,
		validator:  validator,
	}
}
