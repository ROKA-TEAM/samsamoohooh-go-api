package verifier

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/pkg/errors"
	"samsamoohooh-go-api/internal/domain"
	"strings"
	"sync"
)

type Verifier struct {
	validator.Validate
}

func (v *Verifier) Verify(s any) error {
	err := v.Validate.Struct(s)
	if err != nil {
		// Beautify and return the error
		var validationErrors validator.ValidationErrors
		if errors.As(err, &validationErrors) {
			errorMessages := make([]string, len(validationErrors))
			for i, fieldError := range validationErrors {
				errorMessages[i] = fmt.Sprintf("Field '%s' fails '%s' validation", fieldError.Field(), fieldError.Tag())
			}
			return errors.Wrap(domain.ErrValidation, fmt.Sprintf("validation failed:\n%s", strings.Join(errorMessages, "\n")))
		}
		return errors.Wrap(domain.ErrValidation, err.Error())
	}
	return nil
}

var instance Verifier
var once sync.Once

func Get() *Verifier {
	once.Do(func() {
		verifier := Verifier{
			Validate: *validator.New(validator.WithRequiredStructEnabled()),
		}

		instance = verifier
	})

	return &instance
}
