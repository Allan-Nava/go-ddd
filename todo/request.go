package todo

import (
	"fmt"

	"github.com/Allan-Nava/go-ddd/utils"
	"github.com/go-playground/validator"
)

/*
* Copyright © 2022 Allan Nava <>
* Created 05/02/2022
* Updated 14/04/2022
*
 */

type createTodoRequest struct {
	Name string `json:"name" validate:"required,min=4,max=255"`
}

//
var validate = validator.New()

//
func ValidateStruct(request TestValidatorRequest) *utils.ApiError {
	var errorReturn *utils.ApiError
	err := validate.Struct(request)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element utils.ApiError
			//logrus.Error(err)
			element.Message = fmt.Sprint(err) //err.StructNamespace()
			element.Response = "KO"
			errorReturn = &element
		}
	}
	return errorReturn
}
