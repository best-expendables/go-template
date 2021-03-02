package postgresql

import (
	"${REPO_HOST}/${PROJ_NAME}/repository"
	"github.com/best-expendables/common-utils/repository/postgresql"
	"gorm.io/gorm"
)

type ${CAMELIZED_NAME}Repository struct {
	*postgresql.BaseRepo
}

func New${SERVICE_NAME}Repository(db *gorm.DB) repository.${SERVICE_NAME}Repository {
	return &${CAMELIZED_NAME}Repository{
		BaseRepo: postgresql.NewBaseRepo(db),
	}
}
