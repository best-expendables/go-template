package custom

import (
	"${REPO_HOST}/${PROJ_NAME}/constant"
	"gopkg.in/go-playground/validator.v9"
)

func Validate${SERVICE_NAME}Rule(fl validator.FieldLevel) bool {
	input := fl.Field().String()
	return constant.Valid${SERVICE_NAME}Rule[input]
}
