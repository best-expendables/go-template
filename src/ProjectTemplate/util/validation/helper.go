package validation

import (
	"${REPO_HOST}/${PROJ_NAME}/util/validation/custom"
	"reflect"
	"strings"

	"gopkg.in/go-playground/validator.v9"
)

func NewValidator() *validator.Validate {
	validate := validator.New()
	validate.SetTagName("validate")
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]

		if name == "-" {
			return ""
		}

		return name
	})
	registerCustomValidateFunc(validate)
	return validate
}

func registerCustomValidateFunc(validate *validator.Validate) {
	_ = validate.RegisterValidation("${CAMELIZED_NAME}_rule", custom.Validate${SERVICE_NAME}Rule)
}
