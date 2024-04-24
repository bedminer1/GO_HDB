package main

import (
	"github.com/bedminer1/hdb_project/convert"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/2017/records", func(c echo.Context) error {
		return c.JSON(200, convert.CsvToArray(c.QueryParam("town"), c.QueryParam("flat_type"), c.QueryParam("price")))
	})

	e.Logger.Info("Server started on port 8080")
	e.Logger.Fatal(e.Start("localhost:8000"))
}