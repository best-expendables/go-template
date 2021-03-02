package ${FILENAME}

import (
	"${REPO_HOST}/${PROJ_NAME}/model"
	"${REPO_HOST}/${PROJ_NAME}/service/${FILENAME}/dto"
	"context"
)

func (s ${CAMELIZED_NAME}Service) List(ctx context.Context, filter *dto.${SERVICE_NAME}GetListFilter) ([]model.${SERVICE_NAME}, error) {
	var output []model.${SERVICE_NAME}
	err := s.${CAMELIZED_NAME}Repo.Search(ctx, &output, filter)
	if err != nil {
		return nil, err
	}
	return output, nil
}
