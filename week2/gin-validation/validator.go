package main

import (
	"errors"
	"fmt"
	"sync"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
)

var (
	once     sync.Once
	mx       = new(sync.Mutex)
	validate *validator.Validate
)

func New() *validator.Validate {
	once.Do(func() {
		validate = validator.New()
	})

	return validate
}

func TranslateError(err error) error {
	mx.Lock()
	defer mx.Unlock()

	if err == nil {
		return nil

	} else if validate == nil {
		New()
	}

	// Ketika error bukan dari proses validasi data akan langsung dikembalikan tanpa translate
	validationErrors, ok := err.(validator.ValidationErrors)
	if !ok {
		return err

	} else if len(validationErrors) == 0 {
		return nil
	}

	var (
		english  = en.New()
		uni      = ut.New(english, english)
		trans, _ = uni.GetTranslator("en")
	)
	enTranslations.RegisterDefaultTranslations(validate, trans)

	switch e := validationErrors[0]; e.Tag() {
	default:
		return errors.New(e.Translate(trans))

	case "required":
		return fmt.Errorf("%s is a required field", e.Field())
	}
}

func TranslateFormError(err error) map[string]string {
	mx.Lock()
	defer mx.Unlock()

	if err == nil {
		return nil

	} else if validate == nil {
		New()
	}

	// Ketika error bukan dari proses validasi data akan langsung dikembalikan tanpa translate
	validationErrors, ok := err.(validator.ValidationErrors)
	if !ok {
		return map[string]string{"error": err.Error()}

	} else if len(validationErrors) == 0 {
		return nil
	}

	var (
		english  = en.New()
		uni      = ut.New(english, english)
		trans, _ = uni.GetTranslator("en")
	)
	enTranslations.RegisterDefaultTranslations(validate, trans)

	var errorMaps = map[string]string{}

	for _, val := range validationErrors {
		switch val.Tag() {
		default:
			errorMaps[val.Field()] = val.Translate(trans)

		case "required", "required_if":
			errorMaps[val.Field()] = fmt.Sprintf("%s is a required field", val.Field())

		case "email", "boolean":
			errorMaps[val.Field()] = fmt.Sprintf("%s field format must be %s", val.Field(), val.Tag())

		case "eqfield":
			errorMaps[val.Field()] = fmt.Sprintf("field validation for '%s' must equal (%s) '%s'", val.Field(), val.Tag(), val.Param())
		}
	}

	return errorMaps
}
