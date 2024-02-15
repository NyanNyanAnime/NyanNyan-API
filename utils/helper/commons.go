package helper

import (
	"github.com/asaskevich/govalidator"
	"github.com/gorilla/schema"
	"github.com/labstack/echo/v4"
)

func BindFormData(c echo.Context, input interface{}) error {

	if err := c.Bind(input); err != nil {
		return err
	}
	// if err := c.Bind(input); err != nil {
	// 	return err
	// }

	decoder := schema.NewDecoder()
	if err := decoder.Decode(input, c.Request().Form); err != nil {
		return err
	}

	if _, err := govalidator.ValidateStruct(input); err != nil {
		return err
	}

	return nil
}
