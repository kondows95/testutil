package testutil

import (
	"reflect"
)

func Override(base interface{}, override interface{}) interface{} {
	v1 := reflect.Indirect(reflect.ValueOf(base))
	t1 := v1.Type()
	v2 := reflect.Indirect(reflect.ValueOf(override))
	t2 := v2.Type()
	vResult := reflect.New(t1).Elem()
	for i := 0; i < t1.NumField(); i++ {
		ft1 := t1.Field(i)
		fv1 := v1.FieldByName(ft1.Name)
		ft2 := t2.Field(i)
		fv2 := v2.FieldByName(ft2.Name)
		fvResult := vResult.FieldByName(ft1.Name)

		if !fv2.IsZero() {
			fvResult.Set(fv2)
		} else {
			fvResult.Set(fv1)
		}
	}
	return vResult.Interface()
}
