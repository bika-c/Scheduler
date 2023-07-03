package controller

import (
	"scheduler/router/errors"
)

func cast(e error) error {
	switch e.(type) {
	case errors.AccumulateError:
		return e
	case errors.Error:
		return e
	}
	return errors.ErrInvalidJSON
}
