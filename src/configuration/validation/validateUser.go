package validation

import (
	"encoding/json"
	"errors"

	errorhandler "github.com/DreitonWashington/CRUD-GO/src/configuration/errorHandler"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translation "github.com/go-playground/validator/v10/translations/en"
)

var (
	validate = validator.New()
	transl   ut.Translator
)

func Init() {
	if val, ok := binding.Validator.Engine().(*validator.Validate); ok {
		en := en.New()
		unt := ut.New(en, en)
		transl, _ = unt.GetTranslator("en")
		en_translation.RegisterDefaultTranslations(val, transl)
	}
}

func ValidateUserError(validationError error) *errorhandler.Erro {
	var jsonError *json.UnmarshalTypeError
	var jsonValidationError validator.ValidationErrors

	if errors.As(validationError, &jsonError) {
		return errorhandler.NewBadRequestError("Invalid fiel type!")
	} else if errors.As(validationError, &jsonValidationError) {
		errorsCauses := []errorhandler.Causes{}
		for _, e := range validationError.(validator.ValidationErrors) {
			cause := errorhandler.Causes{
				Message: e.Translate(transl),
				Field:   e.Field(),
			}
			errorsCauses = append(errorsCauses, cause)
		}
		return errorhandler.NewBadRequestValidationError("Some field are invalid", errorsCauses)
	} else {
		return errorhandler.NewBadRequestError("Error trying to convert fields")
	}
}
