package validator

import (
	"io"
	"net/http"

	"github.com/labstack/echo/v4"
)

func ValidateRequest(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		body, err := io.ReadAll(c.Request().Body)
		if err != nil || body == nil {
			return c.String(http.StatusBadRequest, "invalid body")
		}

		c.Set("licenseRequest", body)

		return next(c)
	}
}
