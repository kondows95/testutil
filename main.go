package testutil

import (
	"reflect"
)

type Fields map[string]interface{}

func Override(base interface{}, fields Fields) interface{} {
	v := reflect.Indirect(reflect.ValueOf(base))
	t := v.Type()
	vResult := reflect.New(t).Elem()
	for i := 0; i < t.NumField(); i++ {
		ft := t.Field(i)
		fv := v.FieldByName(ft.Name)
		fvResult := vResult.FieldByName(ft.Name)

		fvResult.Set(fv)
		for field, iVal := range fields {
			if field == ft.Name {
				fvResult.Set(reflect.Indirect(reflect.ValueOf(iVal)))
				break
			}
		}
	}
	return vResult.Interface()
}
