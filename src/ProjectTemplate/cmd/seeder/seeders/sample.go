package seeders

import (
	"${REPO_HOST}/${PROJ_NAME}/model"
	utilModel "github.com/best-expendables/common-utils/model"
	"github.com/best-expendables/common-utils/util/uuid_generator"
	"github.com/icrowley/fake"
	"gorm.io/gorm"
)

type ${SERVICE_NAME}s struct {
	*Table
}

func New${SERVICE_NAME}(db *gorm.DB) *${SERVICE_NAME}s {
	return &${SERVICE_NAME}s{NewTable(db)}
}

func (s *${SERVICE_NAME}s) Run() error {
	for i := 0; i < 20; i++ {
		${CAMELIZED_NAME} := model.${SERVICE_NAME}{
			BaseModel: utilModel.BaseModel{
				Id: uuid_generator.NewUUIDV4(),
			},
			Foo: fake.FullName(),
			Bar: fake.FullName(),
		}
		if err := s.Create(&${CAMELIZED_NAME}); err != nil {
			return err
		}

	}
	return nil
}

// ShouldRun Check if seeder should run
func (s *${SERVICE_NAME}s) ShouldRun() bool {
	if s.Count("${CAMELIZED_NAME}s") > 0 {
		return false
	}
	return true
}
