package mw

import (
	"github.com/labstack/echo/v4"
	"log"
	"strings"
)

func RoleCheck(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		val := ctx.Request().Header.Get("User-Role")

		if strings.ToLower(val) == "admin" {
			log.Println("Red button user detected")
		}

		err := next(ctx)

		if err != nil {
			return err
		}
		return nil
	}
}
