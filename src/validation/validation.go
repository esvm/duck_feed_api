package validation

import (
	"strings"

	"github.com/esvm/duck_feed_api/src/exceptions"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	validator "gopkg.in/go-playground/validator.v9"
	en_translations "gopkg.in/go-playground/validator.v9/translations/en"
)

var validate *validator.Validate
var translator ut.Translator

// Validate uses go-validator to validate a struct based on its tags.
// A ValidationError is returned in case of any error.
func Validate(s interface{}) error {
	err := getValidator().Struct(s)

	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			return exceptions.ValidationError{
				Message: err.Error(),
				Details: nil,
			}
		}

		details := map[string]string{}
		for _, err := range err.(validator.ValidationErrors) {
			details[err.Field()] = err.Translate(getTranslator())
		}

		return exceptions.ValidationError{
			Message: err.Error(),
			Details: details,
		}

	}

	return nil
}

// Custom tag created to validate if a field is contained in a set of possible values.
// Currently this only supports strings.
func inValidator(fl validator.FieldLevel) bool {
	value := fl.Field().String()
	param := fl.Param()
	inner := fl.Param()[1 : len(param)-1]
	possible := strings.Split(inner, ";")

	for _, el := range possible {
		if value == el {
			return true
		}
	}

	return false
}

func getValidator() *validator.Validate {
	if validate == nil {
		validate = validator.New()

		en_translations.RegisterDefaultTranslations(validate, getTranslator())
		trans := getTranslator()
		validate.RegisterValidation("in", inValidator)

		validate.RegisterTranslation("in", trans, func(ut ut.Translator) error {
			return ut.Add("in", "This field must be one of {0}", true)
		}, func(ut ut.Translator, fe validator.FieldError) string {
			param := fe.Param()
			inner := fe.Param()[1 : len(param)-1]

			t, _ := ut.T("in", inner)

			return t
		})

		validate.RegisterTranslation("required", trans, func(ut ut.Translator) error {
			return ut.Add("required", "This field is required", true)
		}, func(ut ut.Translator, fe validator.FieldError) string {
			t, _ := ut.T("required")
			return t
		})

		validate.RegisterTranslation("url", trans, func(ut ut.Translator) error {
			return ut.Add("url", "This field must be a valid URL", true)
		}, func(ut ut.Translator, fe validator.FieldError) string {
			t, _ := ut.T("url")
			return t
		})

		validate.RegisterTranslation("max", trans, func(ut ut.Translator) error {
			return ut.Add("max-string", "This field must be a maximum of {0} in length", true)
		}, func(ut ut.Translator, fe validator.FieldError) string {
			t, _ := ut.T("max-string", fe.Param())
			return t
		})

	}

	return validate
}

func getTranslator() ut.Translator {
	if translator == nil {
		en := en.New()
		uni := ut.New(en, en)
		translator, _ = uni.GetTranslator("en")
	}

	return translator
}
