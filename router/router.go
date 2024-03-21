package router

import (
	"com.jtworld.wvproxy/handlers"
	"com.jtworld.wvproxy/validator"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func New() *echo.Echo {
	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.DELETE},
	}))

	// Logger 미들웨어 추가
	e.Use(middleware.Logger())

	v1 := e.Group("/v1")

	v1.POST("/license", handlers.GetLicense, validator.ValidateRequest)

	return e
}
