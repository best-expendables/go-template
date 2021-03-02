package ${FILENAME}

import (
	"${REPO_HOST}/${PROJ_NAME}/model"
	"${REPO_HOST}/${PROJ_NAME}/service/${FILENAME}/dto"
	"github.com/best-expendables/common-utils/service"
	"context"
)

func (s ${CAMELIZED_NAME}Service) Create(ctx context.Context, input dto.${SERVICE_NAME}CreateInput) (*model.${SERVICE_NAME}, error) {
	if err := s.validator.Validate(input); err != nil {
		return nil, service.NewValidationError(err)
	}
	output := input.ConvertToModel()
	err := s.${CAMELIZED_NAME}Repo.Create(ctx, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}
