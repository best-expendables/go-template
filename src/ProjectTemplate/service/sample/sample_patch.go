package ${FILENAME}

import (
	"${REPO_HOST}/${PROJ_NAME}/model"
	"${REPO_HOST}/${PROJ_NAME}/service/${FILENAME}/dto"
	"github.com/best-expendables/common-utils/service"
	"context"
)

func (s ${CAMELIZED_NAME}Service) Patch(ctx context.Context, id string, input dto.${SERVICE_NAME}PatchUpdateInput) (*model.${SERVICE_NAME}, error) {
	if err := s.validator.Validate(input); err != nil {
		return nil, service.NewValidationError(err)
	}
	var output model.${SERVICE_NAME}
	err := s.${CAMELIZED_NAME}Repo.FindByID(ctx, &output, id)
	if err != nil {
		return nil, err
	}
	err = s.${CAMELIZED_NAME}Repo.Update(ctx, &output, input)
	if err != nil {
		return nil, err
	}
	return &output, nil
}
