package main

import (
	"fmt"

	"github.com/samber/lo"
)

func main() {
	fieldValues := []FieldValue{
		{FieldId: 1, Values: []string{"value1", "value2"}},
		{FieldId: 2, Values: []string{"value3", "value4"}},
	}
	datasetId := int32(100)

	// 使用 lo.Map 进行转换
	validateErrors := lo.Map(fieldValues, func(fv FieldValue, _ int) ValidateErr {
		return ValidateErr{
			DatasetId:   datasetId,
			FieldId:     fv.FieldId,
			FieldValues: fv.Values,
		}
	})

	for _, ve := range validateErrors {
		fmt.Printf("DatasetId: %d, FieldId: %d, FieldValues: %v\n", ve.DatasetId, ve.FieldId, ve.FieldValues)
	}

	errs := []ValidateErr{
		{DatasetId: 100, FieldId: 1, FieldValues: []string{"myvalue1", "myvalue2"}},
		{DatasetId: 100, FieldId: 2, FieldValues: []string{"myvalue3", "myvalue4"}},
	}

	values := lo.Map(errs, func(ve ValidateErr, _ int) FieldValue {
		return FieldValue{
			FieldId: ve.FieldId,
			Values:  ve.FieldValues,
		}
	})

	for _, v := range values {
		fmt.Printf("FieldId: %d, Values: %v\n", v.FieldId, v.Values)
	}

}

type FieldValue struct {
	FieldId int32
	Values  []string
}

type ValidateErr struct {
	DatasetId int32
	FieldId   int32

	FieldValues []string
}
