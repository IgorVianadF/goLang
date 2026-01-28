package handler

import (
	"fmt"
	"reflect"
	"time"
)

func errParamIsRequired(name, typ string) error {
	return fmt.Errorf("param: %s (type: %s) is required", name, typ)
}

type SaveDrawingRequest struct {
	Name      string     `json:"name"`
	Price     *float64   `json:"price"`
	Customer  string     `json:"customer"`
	StartDate time.Time  `json:"startDate"`
	EndDate   *time.Time `json:"endDate"`
}

func (r *SaveDrawingRequest) Validate() error {
	v := reflect.ValueOf(r).Elem()
	t := reflect.TypeOf(r).Elem()

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		fieldType := t.Field(i)

		if fieldType.Type.Kind() == reflect.Ptr {
			continue
		}

		if field.IsZero() && fieldType.Name != "Price" {
			return errParamIsRequired(fieldType.Name, fieldType.Type.String())
		}
	}
	return nil
}

func (SaveDrawingRequest) TableName() string {
	return "drawings"
}
