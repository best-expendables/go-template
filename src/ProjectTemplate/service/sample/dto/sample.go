package dto

import (
	"${REPO_HOST}/${PROJ_NAME}/model"
)

type ${SERVICE_NAME}CreateInput struct {
	Foo string `json:"foo" validate:"required"`
	Bar string `json:"bar"`
}

func (p ${SERVICE_NAME}CreateInput) ConvertToModel() *model.${SERVICE_NAME} {
	return &model.${SERVICE_NAME}{
		Foo: p.Foo,
		Bar: p.Bar,
	}
}

type ${SERVICE_NAME}PatchUpdateInput struct {
	Foo *string `json:"foo,omitempty"`
	Bar *string `json:"bar,omitempty"`
}

type ${SERVICE_NAME}PutUpdateInput struct {
	Foo string `json:"foo" validate:"required"`
	Bar string `json:"bar" validate:"required"`
}
