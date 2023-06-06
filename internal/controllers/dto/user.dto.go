package dto

import (
	"fmt"
	"goselflearn/internal/common"
	"strings"

	"github.com/go-playground/validator/v10"
)

type UserRegister struct {
	Name     string `json:"name" validate:"required,min=3,ascii"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password,omitempty" validate:"required,password"`
}

func (t *UserRegister) Validate() *common.ValidationError {
	var errors []*common.ValidationError

	err := common.Validate.Struct(t)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var e common.ValidationError
			e.FailedField = strings.ToLower(err.StructField())
			e.Message = common.GetValidatorErrorMessage(err.Tag(), err.Field(), err.Param())
			errors = append(errors, &e)
		}
	}

	fmt.Println("user validation errors are: ", errors)
	if len(errors) > 0 {
		return errors[0]
	}
	return nil
}

type UserLogin struct {
	Email string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

func (t *UserLogin) Validate() *common.ValidationError {
	var errors []*common.ValidationError

	err := common.Validate.Struct(t)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var e common.ValidationError
			e.FailedField = strings.ToLower(err.StructField())
			e.Message = common.GetValidatorErrorMessage(err.Tag(), err.Field(), err.Param())
			errors = append(errors, &e)
		}
	}

	fmt.Println("user validation errors are: ", errors)
	if len(errors) > 0 {
		return errors[0]
	}
	return nil
}
