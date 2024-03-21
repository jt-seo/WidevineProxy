package handlers

import (
	"log"
	"net/http"

	"com.jtworld.wvproxy/services"
	"github.com/labstack/echo/v4"
)

func GetLicense(c echo.Context) error {
	log.Printf("GetLicense called\n")
	rowRequest := c.Get("licenseRequest").([]byte)

	licenseResponse, err := services.GetLicense(rowRequest)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	return c.String(http.StatusOK, licenseResponse)
}
