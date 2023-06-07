package dto

import (
	"fmt"
	"strings"

	"github.com/ARTM2000/goselflearn/internal/common"

	"github.com/go-playground/validator/v10"
)

type CreatePost struct {
	Title       string `json:"title" validate:"required,min=8" example:"sample title"`
	Description string `json:"description" validate:"required,min=30" example:"sample description for creating a new post with swagger"`
}

func (cp *CreatePost) Validate() *common.ValidationError {
	var errors []*common.ValidationError

	err := common.Validate.Struct(cp)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var e common.ValidationError
			e.FailedField = strings.ToLower(err.StructField())
			e.Message = common.GetValidatorErrorMessage(err.Tag(), err.Field(), err.Param())
			errors = append(errors, &e)
		}
	}

	fmt.Println("create post validation errors are: ", errors)
	if len(errors) > 0 {
		return errors[0]
	}
	return nil
}
