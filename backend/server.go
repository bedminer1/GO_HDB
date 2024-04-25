package main

import (
	"github.com/bedminer1/hdb_project/convert"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/2017/records", func(c echo.Context) error {
		// save query params in options struct
		options := convert.FilterOptions {
			TownFilter: c.QueryParam("town"),
			FlatTypeFilter: c.QueryParam("flat_type"),
			PriceFilter: c.QueryParam("price"),
		}

		// return json with converted csv data
		return c.JSON(200, convert.CsvToJSON(options))
	})

	e.Logger.Print("Listening on port :8080")
	e.Logger.Fatal(e.Start(":8080"))
}