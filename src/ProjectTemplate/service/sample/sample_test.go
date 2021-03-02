package ${FILENAME}_test

import (
	"${REPO_HOST}/${PROJ_NAME}/repository/postgresql"
	"${REPO_HOST}/${PROJ_NAME}/service/${FILENAME}"
	"${REPO_HOST}/${PROJ_NAME}/service/${FILENAME}/dto"
	"${REPO_HOST}/${PROJ_NAME}/util/test"
	"${REPO_HOST}/${PROJ_NAME}/util/validation"
	utilValidator "github.com/best-expendables/common-utils/util/validation"
	"context"
	"github.com/stretchr/testify/suite"
	"log"
	"testing"
)

type ${SERVICE_NAME}TestSuite struct {
	test.MainTestSuite

	${CAMELIZED_NAME}Service ${FILENAME}.${SERVICE_NAME}Service
}

func Test${SERVICE_NAME}TestSuite(s *testing.T) {
	suite.Run(s, new(${SERVICE_NAME}TestSuite))
}

func (s *${SERVICE_NAME}TestSuite) SetupSuite() {
	s.MainTestSuite.SetupSuite()
	validator := validation.NewValidator()
	internalValidator := utilValidator.NewInternalValidator(validator)
	${CAMELIZED_NAME}Repo := postgresql.New${SERVICE_NAME}Repository(s.DB)
	${CAMELIZED_NAME}Service := ${FILENAME}.New${SERVICE_NAME}Service(${CAMELIZED_NAME}Repo, internalValidator)
	s.${CAMELIZED_NAME}Service = ${CAMELIZED_NAME}Service
}

// Test Create ${SERVICE_NAME}
func (s ${SERVICE_NAME}TestSuite) TestCreate() {
	input := dto.${SERVICE_NAME}CreateInput{}
	_, err := s.${CAMELIZED_NAME}Service.Create(context.Background(), input)
	if err != nil {
		log.Fatalf("Fail to create ${CAMELIZED_NAME}: %s", err)
	}
}

// Test Patch ${SERVICE_NAME}
func (s ${SERVICE_NAME}TestSuite) TestPatch() {
	input := dto.${SERVICE_NAME}PatchUpdateInput{}
	id := ""
	_, err := s.${CAMELIZED_NAME}Service.Patch(context.Background(), id, input)
	if err != nil {
		log.Fatalf("Fail to patch update ${CAMELIZED_NAME}: %s", err)
	}
}

// Test Put ${SERVICE_NAME}
func (s ${SERVICE_NAME}TestSuite) TestPut() {
	input := dto.${SERVICE_NAME}PutUpdateInput{}
	id := ""
	_, err := s.${CAMELIZED_NAME}Service.Put(context.Background(), id, input)
	if err != nil {
		log.Fatalf("Fail to put update ${CAMELIZED_NAME}: %s", err)
	}
}

// Test List ${SERVICE_NAME}
func (s ${SERVICE_NAME}TestSuite) TestList() {
	filter := dto.New${SERVICE_NAME}GetListFilter()
	_, err := s.${CAMELIZED_NAME}Service.List(context.Background(), filter)
	if err != nil {
		log.Fatalf("Fail to list ${CAMELIZED_NAME}: %s", err)
	}
}

// Test Delete ${SERVICE_NAME}
func (s ${SERVICE_NAME}TestSuite) TestDelete() {
	id := ""
	err := s.${CAMELIZED_NAME}Service.Delete(context.Background(), id)
	if err != nil {
		log.Fatalf("Fail to delete ${CAMELIZED_NAME}: %s", err)
	}
}
