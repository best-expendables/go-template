package ${FILENAME}

import (
	"${REPO_HOST}/${PROJ_NAME}/model"
	"context"
)

func (s ${CAMELIZED_NAME}Service) Delete(ctx context.Context, id string) error {
	return s.${CAMELIZED_NAME}Repo.DeleteByID(ctx, &model.${SERVICE_NAME}{}, id)
}
