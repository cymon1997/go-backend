package util

import "github.com/cymon1997/go-backend/internal/errors"

func ErrMessage(err error) string {
	errs, ok := err.(*errors.Error)
	if !ok {
		return err.Error()
	}
	return errs.Error()
}

func ErrStatusMessage(err error) string {
	errs, ok := err.(*errors.Error)
	if !ok {
		return "internal server error"
	}
	return errs.Status()
}

func ErrStatus(err error) int {
	errs, ok := err.(*errors.Error)
	if !ok {
		return int(errors.InternalServer)
	}
	return errs.Code()
}

func IsErrNotFound(err error) bool {
	errs, ok := err.(*errors.Error)
	if !ok {
		return false
	}
	return errs.Code() == int(errors.NoDataFound)
}
