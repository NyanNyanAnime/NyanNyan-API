package validation

import (
	"errors"
	"nyannyan/utils/constanta"
)

func CheckDataEmpty(data ...any) error {
	for _, value := range data {
		if value == "" {
			return errors.New(constanta.ERROR_EMPTY)
		}
		if value == 0 {
			return errors.New(constanta.ERROR_EMPTY)
		}
	}
	return nil
}
