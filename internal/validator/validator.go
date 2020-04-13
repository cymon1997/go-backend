package validator

import (
	"fmt"
	"strings"

	"github.com/cymon1997/go-backend/internal/errors"
)

type Validator struct {
	missing []string
	errs    []string
}

func New() *Validator {
	return &Validator{}
}

func (v *Validator) Missing(field string) {
	v.missing = append(v.missing, field)
}

func (v *Validator) Message(message string) {
	v.errs = append(v.errs, message)
}

func (v *Validator) Error() error {
	var message string
	if len(v.missing) > 0 {
		message = fmt.Sprintf("missing field [%s]", strings.Join(v.missing, ","))
	}
	if len(v.errs) > 0 {
		message = fmt.Sprint(message, strings.Join(v.errs, ";"))
	}
	if len(v.missing) == 0 && len(v.errs) == 0 {
		return nil
	}
	return errors.New(errors.InvalidRequest).WithMessage(message)
}
