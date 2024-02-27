package lesson_4

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

func Echo() {
	e := echo.New()

	e.GET("/time", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{
			"time_from_echo": time.Now().Format(time.RFC3339),
		})
	})

	e.Logger.Fatal(e.Start(":8080"))
}
