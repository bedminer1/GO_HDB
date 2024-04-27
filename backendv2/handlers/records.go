package handlers

import (
	"net/http"

	"github.com/bedminer1/echoserver/dbiface"
	"github.com/labstack/echo/v4"
	"gopkg.in/go-playground/validator.v9"
)

// var (
// 	v = validator.New()
// )

// HDBRecord contains data of a HDB resale unit
type HDBRecord struct {

}

// RecordHandler pass in col(reference to mongodb collection) as attribute
type RecordHandler struct {
	Col dbiface.CollectionAPI
}

// ProductValidator class with validate method
type RecordValidator struct {
	validator *validator.Validate
}

// Validate method that validates a product
func (record *RecordValidator) Validate(i interface{}) error {
	return record.validator.Struct(i)
}

func (h *RecordHandler) GetRecords(c echo.Context) error {
	return c.JSON(http.StatusOK, "get")
}