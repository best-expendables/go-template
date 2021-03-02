package ${FILENAME}

import (
	"${REPO_HOST}/${PROJ_NAME}/model"
	"context"
)

func (s ${CAMELIZED_NAME}Service) Get(ctx context.Context, id string, preloadFields ...string) (*model.${SERVICE_NAME}, error) {
	var output model.${SERVICE_NAME}
	err := s.${CAMELIZED_NAME}Repo.FindByID(ctx, &output, id, preloadFields...)
	if err != nil {
		return nil, err
	}
	return &output, nil
}
